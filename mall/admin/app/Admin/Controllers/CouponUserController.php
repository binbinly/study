<?php

namespace App\Admin\Controllers;

use App\Models\CouponUserModel;
use Encore\Admin\Grid;

class CouponUserController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '优惠券详情';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new CouponUserModel());
        $this->disableGridCreate($grid);

        $grid->column('id', 'ID');
        $grid->column('user_id', '用户id');
        $grid->column('coupon_id', '优惠券id');
        $grid->column('is_used', '是否使用')->bool();
        $this->setGridTimeView($grid);

        return $grid;
    }
}
