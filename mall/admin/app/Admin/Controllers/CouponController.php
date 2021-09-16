<?php

namespace App\Admin\Controllers;

use App\Admin\Actions\CouponCreate;
use App\Admin\Actions\Href;
use App\Models\CouponModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class CouponController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '优惠券列表';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new CouponModel());
        $grid->disableExport();

        $grid->actions(function (Grid\Displayers\DropdownActions $actions) {
            $actions->disableView();
            $actions->disableDelete();
            $actions->add(new Href('详情', '/coupon_user'));
            $actions->add(new CouponCreate());
        });

        $grid->column('id', 'ID');
        $grid->column('name', '名称');
        $grid->column('type', '类型')->using(CouponModel::$typeLabel)->label();
        $grid->column('value', '元/折');
        $grid->column('total', '发行数量');
        $grid->column('min_price', '起始价格');
        $grid->column('start_at', '起始时间')->datetime();
        $grid->column('end_at', '结束时间')->datetime();
        $grid->column('desc', '描述');
        $grid->column('is_release', '是否发布')->bool();
        $grid->column('sort', '排序')->editable();
        $this->setGridTimeView($grid);

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new CouponModel());

        $form->text('name', '优惠券名称')->required();
        $form->select('type', '类型')->options(CouponModel::$typeLabel)->required();
        $form->currency('value', '元/折')->symbol('')->required();
        $form->number('total', '发行数量')->required();
        $form->currency('min_price', '起始价格')->required();
        $form->datetime('start_at', '起始时间')->required();
        $form->datetime('end_at', '结束时间')->required();
        $form->text('desc', '描述')->required();
        $form->switch('is_release', '是否发布');
        $form->number('sort', '排序')->default(50);

        return $form;
    }
}
