<?php


namespace App\Models\Goods;


use App\Models\BaseModel;
use Encore\Admin\Traits\ModelTree;

/**
 * 商品分类
 * Class CategoryModel
 * @package App\Models\Goods
 */
class CategoryModel extends BaseModel
{
    use ModelTree;

    protected $table = 'goods_category';

    public function __construct(array $attributes = [])
    {
        parent::__construct($attributes);

        $this->setParentColumn('pid');
        $this->setOrderColumn('sort');
        $this->setTitleColumn('name');
    }

    /**
     * 所有一级分类
     * @return array
     */
    public static function parentAll(){
        return self::query()->where('pid', 0)->pluck('name', 'id')->toArray();
    }

    public static function getAll(){
        return self::query()->pluck('name', 'id')->toArray();
    }

}