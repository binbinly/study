<?php


namespace App\Models;


use App\Admin\Common\Format;

/**
 * 优惠券模型
 * Class CouponModel
 * @package App\Models
 */
class CouponModel extends BaseModel
{
    const TYPE_AMOUNT = 1;
    const TYPE_DISCOUNT = 2;

    public static $typeLabel = [
        self::TYPE_AMOUNT => '满减',
        self::TYPE_DISCOUNT => '折扣'
    ];

    protected $table = 'coupon';

    public function getValueAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setValueAttribute($value)
    {
        $this->attributes['value'] = Format::amountToPenny($value);
    }

    public function getMinPriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setMinPriceAttribute($value)
    {
        $this->attributes['min_price'] = Format::amountToPenny($value);
    }
}