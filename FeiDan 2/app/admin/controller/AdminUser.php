<?php

namespace app\admin\controller;

use think\response\Json;
use think\response\View;
use app\admin\model\AdminUser as AdminUserModel;

class AdminUser extends Base
{
    public function login()
    {
        if($this->request->isPost()) {
            $params = $this->request->param();
            $this->validate($params, [
                'username|用户名' => 'require',
                'password|密码' => 'require',
            ]);

            $adminUser = AdminUserModel::getAdminUserByUsername($params['username']);
            if(empty($adminUser)) return $this->apiError('管理员不存在');
            if($adminUser->status !== 1) return $this->apiError('管理员状态异常');
            if(!$this->passwordVerify($params['password'], $adminUser->password)) return $this->apiError('密码错误');
            unset($adminUser['password']);
            session('admin_user', $adminUser);
            return $this->apiSuccess('登陆成功');
        }
        return $this->view('/adminUser/login');
    }
    
    public function index()
    {
        return $this->view('/user/index');
    }
    
    public function select(): Json
    {
        $adminUser = session('admin_user');
        $page = $this->request->get('page', 1);
        $limit = $this->request->get('limit', 10);
        $model = AdminUserModel::order('id', 'desc');
        
        if($adminUser['id'] != 1) {
            $model = $model->where('id', $adminUser['id']);
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
        $adminUser = session('admin_user');
        if($adminUser['id'] != 1) {
            return $this->apiError('该内容暂无权限');
        }
        
        $params = $this->request->param();
        $this->validate($params, [
            'username|用户名' => 'require',
            'password|密码' => 'require',
        ]);
        
        $params['password'] = $this->passwordHash($params['password']);
        $insertRes = AdminUserModel::insert([
            'username' => $params['username'],
            'password' => $params['password'],
            'create_time' => date('Y-m-d H:i:s'),
            'status' => $params['status'],
        ]);
        if (empty($insertRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }

    public function edit(): Json
    {
        $adminUser = session('admin_user');
        $params = $this->request->param();
        $this->validate($params, [
            'id|数据ID' => 'require',
            'username|用户名' => 'require',
        ]);

        $data = [
            'username' => $params['username'],
            'status' => $params['status'],
            'update_time' => date('Y-m-d H:i:s'),
        ];
        
        if(!empty($params['password'])) {
            $data['password'] = $this->passwordHash($params['password']);
        }
        
        if($adminUser['id'] != 1 && $params['id'] != $adminUser['id']) {
            return $this->apiError('只能修改自己的账户');
        }
        
        $updateRes = AdminUserModel::where('id', $params['id'])->update($data);
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }

    public function del(): Json
    {
        $adminUser = session('admin_user');
        if($adminUser['id'] != 1) {
            return $this->apiError('该内容暂无权限');
        }
        
        $params = $this->request->param();
        if(empty($params)) return $this->apiError('删除数据ID错误');
        if($params['id'] === 1) return $this->apiError('不能删除根用户');
        $updateRes = AdminUserModel::where('id', $params['id'])->delete();
        if (empty($updateRes)) return $this->apiError('操作失败, 请重试');
        return $this->apiSuccess('操作成功');
    }
    
    public function logout()
    {
        $adminUser = session('admin_user');
        session('admin_user', null);
        return redirect('/admin');
        // return $this->apiSuccess('操作成功');
    }
}