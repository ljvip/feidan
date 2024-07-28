<?php

namespace app\admin\controller;

class Index extends Base
{
    public function index()
    {
        return redirect('/admin/platform');
        // return $this->view('/platform/index');
    }
}