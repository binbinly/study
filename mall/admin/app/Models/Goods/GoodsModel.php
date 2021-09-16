<?php


namespace App\Models\Goods;

use App\Admin\Common\Format;
use App\Models\BaseModel;
use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * 商品模型
 * Class GoodsModel
 * @package App\Models
 */
class GoodsModel extends BaseModel
{
    use SoftDeletes;

    const STATUS_INIT = 0;
    const STATUS_ONLINE = 1;
    const STATUS_OFFLINE = 2;

    public static $statusLabel = [
        self::STATUS_INIT => '草稿',
        self::STATUS_ONLINE => '上架',
        self::STATUS_OFFLINE => '下架'
    ];

    protected $table = 'goods';

    public function getPriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setPriceAttribute($value)
    {
        $this->attributes['price'] = Format::amountToPenny($value);
    }

    public function getOriginalPriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setOriginalPriceAttribute($value)
    {
        $this->attributes['original_price'] = Format::amountToPenny($value);
    }
}