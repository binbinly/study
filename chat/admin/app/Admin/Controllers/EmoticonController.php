<?php

namespace App\Admin\Controllers;

use App\Admin\Extensions\Tools\GridView;
use App\Models\EmoticonModel;
use Encore\Admin\Controllers\AdminController;
use Encore\Admin\Form;
use Encore\Admin\Grid;
use Illuminate\Support\Facades\Request;

class EmoticonController extends AdminController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'EmoticonModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new EmoticonModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('category', '分类');
        $grid->column('name', '名称');
        $grid->column('url', '资源地址')->image();
        $grid->column('created_at', '创建时间');
        $grid->column('updated_at', '更新时间');

        $grid->tools(function ($tools) {
            $tools->append(new GridView());
        });

        if (Request::get('view') !== 'table') {
            $grid->setView('admin.grid.card');
        }
        $grid->filter(function (Grid\Filter $filter) {
            $filter->disableIdFilter();
            $filter->equal('category', '分类');
        });

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new EmoticonModel());

        $form->text('category', '分类');
        $form->text('name', '名称');
        $form->image('url', '资源地址');

        return $form;
    }
}
