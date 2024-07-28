<?php

use think\facade\Route;
use app\admin\middleware\Auth;


// 登录
Route::any('/admin_user/login', 'adminUser/login');

Route::group(function () {
    Route::get('/', 'index/index');
    
    Route::get('/platform', 'platform/index');
    Route::get('/platform/select', 'platform/select');
    Route::post('/platform/add', 'platform/add');
    Route::post('/platform/edit', 'platform/edit');
    Route::post('/platform/del', 'platform/del');
    Route::post('/platform/getToken', 'platform/getToken');
    Route::post('/platform/getUserInfo', 'platform/getUserInfo');
    Route::post('/platform/getList', 'platform/getList');
    Route::get('/platform/select_log', 'platform/select_log');
    Route::post('/platform/del_log', 'platform/del_log');
    Route::get('/platform/log', 'platform/log');
    
    Route::get('/user', 'adminUser/index');
    Route::get('/user/select', 'adminUser/select');
    Route::post('/user/add', 'adminUser/add');
    Route::post('/user/edit', 'adminUser/edit');
    Route::post('/user/del', 'adminUser/del');
    Route::get('/user/logout', 'adminUser/logout');
})->middleware([
    Auth::class
]);