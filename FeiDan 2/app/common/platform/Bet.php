<?php

namespace app\common\platform;

use app\common\Http;
use Exception;

class Bet extends Base
{
    public static function bet(array $data = [])
    {
        $data = json_encode($data, 256);
        $response = Http::post(static::getApi('bet'), $data, [
            'Content-Type: application/json; charset=utf-8',
        ]);
        $response = json_decode($response, true);
        if ($response['code'] !== 200) throw new Exception($response['msg'] ?? '服务器处理错误, 请重试');
        return $response['data'];
    }
}