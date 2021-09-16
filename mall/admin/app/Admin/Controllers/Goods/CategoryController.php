<?php


namespace App\Admin\Controllers\Goods;


use App\Admin\Controllers\BaseController;
use App\Models\Goods\CategoryModel;
use Encore\Admin\Form;
use Encore\Admin\Tree;
use Encore\Admin\Layout\Content;

class CategoryController extends BaseController
{
    public function index(Content $content)
    {
        $tree = new Tree(new CategoryModel);

        return $content
            ->header('商品分类树')
            ->body($tree);
    }

    protected function form()
    {
        $form = new Form(new CategoryModel());

        $form->select('pid', '父ID')->options(CategoryModel::selectOptions(null, '所属分类'))->required();
        $form->text('name', '分类名')->required();
        $form->textarea('desc', '描述');
        $form->number('sort', '排序值')->default(50)->help('值越大越靠前');

        return $form;
    }
}