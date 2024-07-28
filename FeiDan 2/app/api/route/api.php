<?php

use think\facade\Route;

Route::any('/outFly', 'feidan/notify');
Route::get('/auto_login', 'feidan/autoLogin');