<?php


namespace App\Admin\Forms;


use App\Admin\Services\ConfigService;
use App\Models\AppConfigModel;
use App\Models\Goods\CategoryModel;
use Encore\Admin\Form\NestedForm;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Exception;
use Throwable;

class Pay extends Base
{
    public $title = '支付配置';

    /**
     * @param Request $request
     * @return RedirectResponse
     * @throws Throwable
     */
    public function handle(Request $request)
    {
        parent::handle($request);

        $data = $request->all();
        $list = [];
        foreach($data[AppConfigModel::KEY_PAY_LIST] as $item) {
            if ($item['_remove_'] == 1) {
                continue;
            }
            $item['id'] = intval($item['id']);
            $list[] = $item;
        }

        ConfigService::save([AppConfigModel::KEY_PAY_LIST => json_encode($list)]);

        return $this->success();
    }

    /**
     * Build a form here.
     */
    public function form()
    {
        $this->table(AppConfigModel::KEY_PAY_LIST, '以太坊', function (NestedForm $form){
            $form->number('id', '主键')->help('主键请不要重复');
            $form->text('name', '币');
            $form->text('address', '地址');
        });
    }

    /**
     * @return array|mixed
     * @throws Exception
     */
    public function data()
    {
        return [AppConfigModel::KEY_PAY_LIST => AppConfigModel::init(AppConfigModel::KEY_PAY_LIST, AppConfigModel::TYPE_JSON)];
    }
}