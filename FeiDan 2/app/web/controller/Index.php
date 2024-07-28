<?php

namespace app\web\controller;

use app\BaseController;

class Index extends BaseController
{
    public function index()
    {
        return app()->version();
    }
}
