<?php

namespace app\common\platform;

class Base
{
    private static string $BASE_URL = "http://127.0.0.1:8000";
    private static array $API = [
        'getToken' => '/api/v1/account/getToken',
        'getUserInfo' => '/api/v1/account/getUserInfo',
        'bet' => '/api/v1/bet',
        'getList' => '/api/v1/getList'
    ];

    public static function getApi(string $key = ''): string
    {
        return static::$BASE_URL . static::$API[$key];
    }
}