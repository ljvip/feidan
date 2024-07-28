<?php

namespace app\api\controller;

use app\admin\model\Platform;
use app\common\platform\Account;
use app\common\platform\Bet;
use Exception;
use think\facade\Db;
use think\facade\Log;

class Feidan extends Base
{
    public function api_login($val = [])
    {
        try {
            $token = Account::getToken([
                'url' => $val['url'],
                'username' => $val['username'],
                'password' => $val['password'],
            ]);
            
            Platform::where('id', $val['id'])->update([
                'update_time' => date('Y-m-d H:i:s'),
                'online' => 1,
                'offline_reason' => '',
                'token' => $token,
            ]);
            
            return $token;
        } catch (Exception $e) {
            $data = [
                'online' => 0,
                'update_time' => date('Y-m-d H:i:s'),
                'offline_reason' => $e->getMessage(),
            ];
            
            if($e->getMessage() == '账号或密码错误。') {
                $data['status'] = 0;
            }
            Platform::where('id', $val['id'])->update($data);
        }
        
        return false;
    }
    
    public function api_get_user_info($val = [])
    {
        try {
            $userInfo = Account::getUserInfo([
                'url' => $val['url'],
                'token' => $val['token'],
            ]);
            if(isset($userInfo['accounts'][0]['balance'])) {
                $info = [
                    'balance' => $userInfo['accounts'][0]['balance'] ?? '0.00',
                    'betting' => $userInfo['accounts'][0]['betting'] ?? '0.00',
                    'result' => $userInfo['accounts'][0]['result'] ?? '0.00',
                ];
                Platform::where('id', $val['id'])->update([
                    'update_time' => date('Y-m-d H:i:s'),
                    'balance' => $info['balance'],
                    'betting' => $info['betting'],
                    'result' => $info['result'],
                ]);
                return $userInfo;
            }
            
            return false;
        } catch (Exception $e) {
            return false;
        }
    }
    
    public function autoLogin()
    {
        $platform = Platform::where('status', 1)->order('polling', 'asc')->select()->toArray();
        $account = [];
        $countLogin = 0;
        
        // 遍历用户列表
        foreach ($platform as $key => $val) {
            // 获取用户信息
            $userInfo = $this->api_get_user_info($val);
            if(empty($userInfo)) {
                // 获取用户信息失败，重新登陆
                $token = $this->api_login($val);
                if(!empty($token)) {
                    $val['token'] = $token;
                    $userInfo = $this->api_get_user_info($val);
                    $countLogin++;
                }
            }
            
            if(isset($userInfo) && !empty($userInfo)) $account[] = $val;
        }
        
        $data = '';
        foreach ($account as $k => $v) {
            $data .= '平台名称: <a href="' . $v['url'] . '">' . $v['platform_name'] . '</a>' . '<br>用户名: ' . $v['username'] . '<hr>';
        }
        return '全部账户: ' . count($platform) . '<br>本次登录: ' . $countLogin . '<br>在线账户: ' . count($account) . '<br><br>' . $data;
    }
    
    public function notify()
    {
        $params = $this->request->param();
        
        if (empty($params) || empty($params['data'])) return json([
            'state' => 0,
            'success' => [],
            'failure' => []
        ]);
        
        // Token 不正确
        if(empty($params['token']) || $params['token'] != 'f3785tg3b48f237fg8243yt5') return json([
            'state' => 0,
            'success' => [],
            'failure' => []
        ]);
        
        Log::error('接收到API参数' . json_encode($params, 256));
        
        $failData = [
            'state' => 0,
            // 'error' => 'No account',
            'success' => [],
            'failure' => $params['data'] ?? []
        ];
        
        // 所有可用账户
        $platform = Platform::alias('p')
        ->join('admin_user a', 'p.admin_user_id = a.id')
        ->field('a.username as a_username,p.*')->where('p.status', 1)->order('p.polling', 'asc')->select()->toArray();
        
        // 计算总金额
        $maxAmount = 0;
        foreach ($params['data'] as $k => $v) {
            foreach ($v['list'] as $kk => $vv) {
                $maxAmount += $vv['amount'];
            }
        }
        
        $cePlatform = []; // 有代理的账户
        if(!empty($params['ce'])) {
            $ce = explode(",", $params['ce']); // post 指定的代理
            foreach($platform as $k => $v) {
                $maxAmountNew = $maxAmount * ($v['redouble'] / 100); // 乘加倍金额
                $sysCe = explode(",", $v['ce']); // 设置的代理
                foreach ($ce as $kk => $vv) {
                    // post 指定的代理 === 后台设置的代理用户名 && 当前余额 > 总金额时
                    if($vv === $v['a_username'] && $v['balance'] >= $maxAmountNew) {
                        // 没有设置加倍 或者金额在止盈止损范围内
                        if($v['redouble'] == 100 || ($v['result'] < $v['redouble_profit'] && $v['result'] > -$v['redouble_loss'])) {
                            $cePlatform[] = $v;
                        }
                        continue 2;
                    }
                    
                    // post 指定的代理 === 后台设置的代理 && 当前余额 > 总金额时
                    foreach ($sysCe as $kkk => $vvv) {
                        if($vvv === $vv && $v['balance'] >= $maxAmountNew) {
                            // 没有设置加倍 或者金额在止盈止损范围内
                            if($v['redouble'] == 100 || ($v['result'] < $v['redouble_profit'] && $v['result'] > -$v['redouble_loss']) ) {
                                $cePlatform[] = $v;
                            }
                            continue 3;
                        }
                    }
                }
            }
        }
        
        // 没有找到代理账户
        if(empty($cePlatform)) {
            Log::error('没有找到代理账户-返回API结果' . json_encode($failData, 256));
            
            // 添加错误信息参数
            foreach ($failData['failure'] as $dataK => $dataV) {
                foreach ($dataV['list'] as $listK => $listV) {
                    $failData['failure'][$dataK]['list'][$listK]['error'] = '无可用账户';
                }
            }
            
            return json($failData);
        }
        
        // 按使用次数重新排序
        $polling = array_column($cePlatform, 'polling');
        array_multisort($polling, SORT_ASC, $cePlatform);
        
        // 优先使用次数最少的账户发送请求
        $result = null;
        $currentCe = null;
        $logId = null; // 飞单记录ID
        
        foreach ($cePlatform as $k => $v) {
            $userInfo = $this->api_get_user_info($v);
            // 检查token是否有效 如果有效发送请求
            if(!empty($userInfo)) {
                $params['url'] = $v['url'];
                $params['token'] = $v['token'];
                
                try {
                    // 插入飞单记录
                    $logId = $this->addLog([
                        'admin_user_id' => $v['admin_user_id'],
                        'url' => $v['url'],
                        'username' => $v['username'],
                        'create_time' => date("Y-m-d H:i:s"),
                        'send' => json_encode($params, 256)
                    ]);
                    
                    Platform::where('id', $v['id'])->inc('polling', 1)->update();
                    
                    // 加倍
                    if($v['redouble'] != 100) {
                        foreach ($params['data'] as $dataK => $dataV) {
                            foreach ($dataV['list'] as $listK => $listV) {
                                $params['data'][$dataK]['list'][$listK]['amount'] = $params['data'][$dataK]['list'][$listK]['amount'] * ($v['redouble'] / 100);
                            }
                        }
                    }
                    
                    Log::error('开始飞单' . json_encode($params, 256));
                    $result = Bet::bet($params);
                    if(!empty($result)) {
                        $currentCe = $v;
                        break;
                    }
                } catch (Exception $e) {
                    // 如果报错 继续下一个
                    continue;
                }
            }
        }
        
        // 如果所有帐户都失败
        if (empty($result) || empty($currentCe) || empty($logId)) {
            // $failData['error'] = 'Flight ticket failure';
            // 添加错误信息参数
            foreach ($failData['failure'] as $dataK => $dataV) {
                foreach ($dataV['list'] as $listK => $listV) {
                    $failData['failure'][$dataK]['list'][$listK]['error'] = '飞单失败';
                }
            }
            
            Log::error('飞单失败-返回API结果' . json_encode($failData, 256));
            return json($failData);
        }
        
        // 更新飞单记录
        $this->editLog($logId, [
            'input' => json_encode($result, 256),
            'update_time' => date('Y-m-d H:i:s'),
        ]);
        
        Log::error('返回API结果' . json_encode($result['data'], 256));
        return json($result['data']);
    }
    
    public function addLog($data)
    {
        Log::error('添加日志' . json_encode($data, 256));
        $for = true;
        while ($for) {
            try {
                $res = Db::table('platform_log')->insertGetId($data);
                if(!empty($res)) {
                    break;
                }
            } catch (Exception $e) {}
        }
        
        return $res;
    }
    
    public function editLog($logId, $data)
    {
        Log::error('更新日志' . $logId . json_encode($data, 256));
        $for = true;
        while ($for) {
            try {
                $res = Db::table('platform_log')->where('id', $logId)->update($data);
                if(!empty($res)) {
                    break;
                }
            } catch (Exception $e) {}
        }
        
        return $res;
    }
    
    public function isFullWordMatch($text, $word) {
        $pattern = "/\b" . preg_quote($word) . "\b/";
        $matches = preg_match($pattern, $text);
        return $matches ? true : false;
    }
}
