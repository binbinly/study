<?php

namespace App\Admin\Actions;

use App\Models\CouponUserModel;
use Encore\Admin\Actions\RowAction;
use Illuminate\Database\Eloquent\Model;

/**
 * 生成优惠券
 * Class CouponCreate
 * @package App\Admin\Actions
 */
class CouponCreate extends RowAction
{
    public $name = '生成优惠券';

    public function handle(Model $model)
    {
        //已经生成数量
        $count = CouponUserModel::query()->where('coupon_id', $model->id)->count();
        //待生成数量
        $num = $model->total - $count;
        if ($num > 0) {
            $now = time();
            $data = array_fill(0, $num, ['coupon_id' =>$model->id, 'user_id' => 0, 'created_at' => $now, 'updated_at' => $now]);
            CouponUserModel::insert($data);
        }
        return $this->response()->success('操作成功')->refresh();
    }

    public function dialog()
    {
        $this->confirm('确定生成？');
    }

}