<?php
namespace app\admin\model;

use think\db\exception\DataNotFoundException;
use think\db\exception\DbException;
use think\db\exception\ModelNotFoundException;
use think\model;

class AdminUser extends Model
{
    /**
     * 根据用户名获取信息
     * @param string $username
     * @return array|mixed
     * @throws DataNotFoundException
     * @throws DbException
     * @throws ModelNotFoundException
     */
    public static function getAdminUserByUsername(string $username): mixed
    {
        return self::where('username', $username)->find();
    }
}