<?php
declare (strict_types=1);

namespace app\admin\middleware;

use Closure;

class Auth
{
    public function handle($request, Closure $next)
    {
        if (empty(session('admin_user'))) {
            return redirect('/admin/admin_user/login');
        }

        return $next($request);
    }
}
