<?php

use Illuminate\Routing\Router;

Admin::routes();

Route::group([
    'prefix'        => config('admin.route.prefix'),
    'namespace'     => config('admin.route.namespace'),
    'middleware'    => config('admin.route.middleware'),
    'as'            => config('admin.route.prefix') . '.',
], function (Router $router) {

    $router->get('/', 'HomeController@index')->name('home');
    $router->resource('goods', 'Goods\GoodsController')->names('商品管理');
    $router->get('app_config', 'AppConfigController@index')->name('配置管理');
    $router->resource('app_setting', 'AppSettingController')->names('页面设置');
    $router->resource('goods_category', 'Goods\CategoryController')->names('商品分类');
    $router->resource('app_notice', 'AppNoticeController')->names('公告管理');
    $router->resource('area', 'AreaController')->names('地址管理');
    $router->resource('coupon', 'CouponController')->names('优惠券管理');
    $router->resource('coupon_user', 'CouponUserController')->names('优惠券详情');

});
