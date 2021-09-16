<?php

namespace App\Admin\Services;

use App\Models\AppConfigModel;
use Throwable;

/**
 * 配置相关
 * Class ConfigService
 * @package AdminCommon\Services
 */
class ConfigService
{
    /**
     * 配置修改保存
     * @param $data
     * @return bool
     * @throws Throwable
     */
    public static function save($data)
    {
        $ret = false;
        foreach ($data as $name => $value) {
            $ret = self::saveOnw($name, $value);
        }
        return $ret;
    }

    /**
     * 添加一条配置
     * @param string $name
     * @param $value
     * @return bool
     */
    public static function saveOnw(string $name, $value)
    {
        $model = AppConfigModel::query()->where('name', $name)->first();
        if ($model instanceof AppConfigModel) {
            if ($model->value == $value) return true;
        } else {
            $model = new AppConfigModel();
            $model->name = $name;
        }
        $model->value = $value;
        return $model->save();
    }
}