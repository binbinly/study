<?php

namespace App\Admin\Controllers;

use App\Models\AppNoticeModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class AppNoticeController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'AppNoticeModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new AppNoticeModel());
        $grid->disableExport();
        $grid->model()->orderBy('id', 'desc');

        $grid->column('id', 'ID');
        $grid->column('title', '标题');
        $grid->column('content', '内容');
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
        $form = new Form(new AppNoticeModel());

        $form->text('title', '标题');
        $form->textarea('content', '内容');

        return $form;
    }
}
