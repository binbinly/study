<?php

namespace App\Admin\Controllers;

use App\Admin\Common\Search;
use App\Models\AreaModel;
use Encore\Admin\Grid;

class AreaController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'AreaModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new AreaModel());
        $this->disableGridExport($grid);

        $grid->column('id', 'ID');
        $grid->column('level', '层级');
        $grid->column('parent_code', '上级');
        $grid->column('area_code', '行政代码');
        $grid->column('zip_code', '邮编');
        $grid->column('city_code', '区号');
        $grid->column('name', '名称');
        $grid->column('short_name', '简称');
        $grid->column('merger_name', '全称');
        $grid->column('pinyin', '拼音');
        $grid->column('lng', '经度');
        $grid->column('lat', '纬度');

        return $grid;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('area_code', '行政代码');
        $filter->equal('parent_code', '上级行政代码');
    }
}
