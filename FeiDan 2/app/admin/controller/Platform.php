<?php

namespace app\admin\controller;

use app\common\platform\Account;
use Exception;
use think\facade\Db;
use think\response\Json;
use think\response\View;
use app\admin\model\Platform as PlatformModel;

class Platform extends Base
{
    public function index(): View
    {
        return $this->view('/platform/index');
    }

    public function log()
    {
        return $this->view('/platform/log');
    }

    public function select_log(): Json
    {
        $page = $this->request->get('page', 1);
        $limit = $this->request->get('limit', 10);
        $model = Db::table('platform_log')->order('id', 'desc');
        
        $adminUser = session('admin_user');
        if($adminUser['id'] != 1) {
            $model = $model->where('admin_user_id', $adminUser['id']);
        }
        
        $count = $model->count();
        $data = $model->page($page, $limit)->select()->toArray();

        return $this->apiSuccess('查询成功', [
            'count' => $count,
            'data' => $data
        ]);
    }
    
    public function del_log(): Json
    {
        $adminUser = session('admin_user');
        if($adminUser['id'] != 1) {
            return $this->apiError('该内容暂无权限');
        }
        
        $params = $this->request->param();
        if(empty($params)) return $this->apiError('删除数据ID错误');
        $updateRes = Db::table('platform_log')
                        ->where('id', '>=', $params['start_id'])
                        ->where('id', '<=', $params['end_id'])
                        ->delete();
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }


    public function select(): Json
    {
        $page = $this->request->get('page', 1);
        $limit = $this->request->get('limit', 10);
        $model = PlatformModel::alias('p')
        ->join('admin_user a', 'p.admin_user_id = a.id')
        ->field('a.username as a_username,p.*')->order('p.id', 'desc');
        
        $adminUser = session('admin_user');
        if($adminUser['id'] != 1) {
            $model = $model->where('p.admin_user_id', $adminUser['id']);
        }else{
            return $this->apiSuccess('ok', [
                'count'=> 0,
                'data' => []
            ]);
        }
        
        $count = $model->count();
        $data = $model->page($page, $limit)->select()->toArray();

        return $this->apiSuccess('查询成功', [
            'count' => $count,
            'data' => $data
        ]);
    }

    public function add(): Json
    {
        $params = $this->request->param();
        $this->validate($params, [
            'platform_name|平台名称' => 'require',
            'url|平台地址' => 'require',
            'username|平台账户' => 'require',
            'password|平台密码' => 'require',
        ]);
        
        $adminUser = session('admin_user');
        $insertRes = PlatformModel::insert([
            'admin_user_id' => $adminUser['id'],
            'platform_name' => $params['platform_name'],
            'url' => $params['url'],
            'username' => $params['username'],
            'password' => $params['password'],
            'ce' => $params['ce'] ?? '',
            'auto_login' => $params['auto_login'],
            'redouble' => $params['redouble'],
            'redouble_profit' => $params['redouble_profit'],
            'redouble_loss' => $params['redouble_loss'],
            'create_time' => date('Y-m-d H:i:s'),
            'status' => $params['status'],
        ]);
        if (empty($insertRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }

    public function edit(): Json
    {
        $params = $this->request->param();
        $this->validate($params, [
            'id|数据ID' => 'require',
            'platform_name|平台名称' => 'require',
            'url|平台地址' => 'require',
            'username|平台账户' => 'require',
            'password|平台密码' => 'require',
        ]);

        $updateRes = PlatformModel::where('id', $params['id'])->update([
            'platform_name' => $params['platform_name'],
            'url' => $params['url'],
            'username' => $params['username'],
            'password' => $params['password'],
            'ce' => $params['ce'] ?? '',
            'auto_login' => $params['auto_login'],
            'redouble' => $params['redouble'],
            'redouble_profit' => $params['redouble_profit'],
            'redouble_loss' => $params['redouble_loss'],
            'status' => $params['status'],
            'update_time' => date('Y-m-d H:i:s'),
        ]);
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }

    public function del(): Json
    {
        $params = $this->request->param();
        if(empty($params)) return $this->apiError('删除数据ID错误');
        $updateRes = PlatformModel::where('id', $params['id'])->delete();
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }

    public function getList()
    {
        $params = $this->request->param();
        $platform = PlatformModel::find($params['id']);
        $result = Account::getList([
            'url' => $platform['url'],
            'token' => $platform['token']
        ]);

        return $this->apiSuccess('获取成功', $result['result']);
    }

    public function getToken(): Json
    {
        $params = $this->request->param();
        $this->validate($params, [
            'id|数据ID' => 'require'
        ]);

        $data = PlatformModel::find($params['id'])->toArray();
        if (empty($data)) return $this->apiError('暂未查询到该数据, 请重试');

        // 获取token
        try {
            $token = Account::getToken([
                'url' => $data['url'],
                'username' => $data['username'],
                'password' => $data['password'],
            ]);
        } catch (Exception $e) {
            PlatformModel::where('id', $params['id'])->update([
                'online' => 0,
                'update_time' => date('Y-m-d H:i:s'),
                'offline_reason' => $e->getMessage(),
            ]);
            return $this->apiError($e->getMessage());
        }

        // 获取余额信息
        $userInfo = Account::getUserInfo([
            'url' => $data['url'],
            'token' => $token,
        ]);

        $info = [
            'balance' => $userInfo['accounts'][0]['balance'] ?? '0.00',
            'betting' => $userInfo['accounts'][0]['betting'] ?? '0.00',
            'result' => $userInfo['accounts'][0]['result'] ?? '0.00',
        ];

        // 更新账号信息
        $updateRes = PlatformModel::where('id', $params['id'])->update([
            'token' => $token,
            'online' => 1,
            'offline_reason' => '',
            'update_time' => date('Y-m-d H:i:s'),
            'balance' => $info['balance'],
            'betting' => $info['betting'],
            'result' => $info['result'],
        ]);
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('登录成功');
    }

    public function getUserInfo(): Json
    {
        $params = $this->request->param();
        $this->validate($params, [
            'id|数据ID' => 'require'
        ]);

        $data = PlatformModel::find($params['id']);
        if (empty($data)) return $this->apiError('暂未查询到该数据, 请重试');
        if (empty($data['token']) || $data['online'] === 0) return $this->apiError('该账户暂未登录, 请先登录');

        // 获取余额信息
        try {
            $userInfo = Account::getUserInfo([
                'url' => $data['url'],
                'token' => $data['token'],
            ]);
        } catch (Exception $e) {
            PlatformModel::where('id', $params['id'])->update([
                'online' => 0,
                'update_time' => date('Y-m-d H:i:s'),
                'offline_reason' => $e->getMessage(),
            ]);
            return $this->apiError($e->getMessage());
        }

        $info = [
            'balance' => $userInfo['accounts'][0]['balance'] ?? '0.00',
            'betting' => $userInfo['accounts'][0]['betting'] ?? '0.00',
            'result' => $userInfo['accounts'][0]['result'] ?? '0.00',
        ];
        // 更新账号信息
        $updateRes = PlatformModel::where('id', $params['id'])->update([
            'update_time' => date('Y-m-d H:i:s'),
            'balance' => $info['balance'],
            'betting' => $info['betting'],
            'result' => $info['result'],
        ]);
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }
}