<?php

namespace app\admin\controller;

use app\BaseController;
use think\response\Json;
use think\response\View;

class Base extends BaseController
{
    public function view(...$args): View
    {
        return view(...$args);
    }

    public function apiResult(int $code = 200, string $msg = 'Success', array $data = [], int $statusCode = 200): Json
    {
        return json([
            'code' => $code,
            'msg' => $msg,
            'data' => $data
        ], $statusCode);
    }

    public function apiSuccess(string $msg = 'Success', array $data = []): Json
    {
        return $this->apiResult(200, $msg, $data);
    }

    public function apiError(string $msg = 'Error', array $data = []): Json
    {
        return $this->apiResult(500, $msg, $data);
    }

    /**
     * 密码哈希
     * @param string $password
     * @return string
     */
    public function passwordHash(string $password): string
    {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    /**
     * 验证密码哈希
     * @param string $password
     * @param string $hash
     * @return bool
     */
    public function passwordVerify(string $password, string $hash): bool
    {
        return password_verify($password, $hash);
    }
}