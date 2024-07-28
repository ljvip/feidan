<?php

namespace app\common\platform;

use app\common\Http;
use Exception;

class Account extends Base
{
    /**
     * 登录-获取token
     * @param array $data
     * @return string
     * @throws Exception
     */
    public static function getToken(array $data = []): string
    {
        $response = Http::post(static::getApi('getToken'), [
            'url' => $data['url'],
            'username' => $data['username'],
            'password' => $data['password'],
        ]);
        $response = json_decode($response, true);
        if ($response['code'] !== 200) throw new Exception($response['msg'] ?? '服务器处理错误, 请重试');
        return $response['data']['token'];
    }

    public static function getUserInfo(array $data = []): array
    {
        $response = Http::post(static::getApi('getUserInfo'), [
            'url' => $data['url'],
            'token' => $data['token'],
        ]);
        $response = json_decode($response, true);
        if ($response['code'] !== 200) throw new Exception($response['msg'] ?? '服务器处理错误, 请重试');
        return $response['data'];
    }

    public static function getList(array $data = [])
    {
        $response = Http::post(static::getApi('getList'), [
            'url' => $data['url'],
            'token' => $data['token'],
        ]);
        $response = json_decode($response, true);
        if ($response['code'] !== 200) throw new Exception($response['msg'] ?? '服务器处理错误, 请重试');
        return $response['data'];
    }
}