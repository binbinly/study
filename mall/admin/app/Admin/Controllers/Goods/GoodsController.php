<?php

namespace App\Admin\Controllers\Goods;

use App\Admin\Common\Search;
use App\Admin\Controllers\BaseController;
use App\Models\Goods\CategoryModel;
use App\Models\Goods\GoodsModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class GoodsController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'GoodsModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new GoodsModel());

        $grid->column('id', 'ID');
        $grid->column('title', '标题');
        $grid->column('cat_id', '分类')->using(CategoryModel::getAll())->filter();;
        $grid->column('cover', '封面')->image();
        $grid->column('price', '市场价');
        $grid->column('original_price', '原价');
        $grid->column('stock', '库存');
        $grid->column('sku_many', '多规格')->bool();
        $grid->column('status', '状态')->using(GoodsModel::$statusLabel);
        $grid->column('sale_count', '销量');
        $grid->column('review_count', '评论数');
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
        $form = new Form(new GoodsModel());

        $form->text('title', '标题');
        $form->number('cat_id', '分类')->options(CategoryModel::selectOptions(null, '所属分类'))->required();
        $form->image('cover', '封面');
        $form->currency('price', '市场价');
        $form->currency('original_price', '原价');
        $form->textarea('intro', '简介');
        $form->text('unit', '单位');
        $form->number('stock', '库存');
        $form->switch('sku_many', '多规格');
        $form->switch('status', '状态')->options(GoodsModel::$statusLabel);
        $form->number('discount', '折扣');
        $form->number('sale_count', '销量');
        $form->number('review_count', '评论数');
        $form->number('sort', '排序')->default(50);

        return $form;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('id', '商品ID');
        // 范围过滤器，调用模型的`onlyTrashed`方法，查询出被软删除的数据。
        $filter->scope('trashed', '回收站')->onlyTrashed();
    }
}
