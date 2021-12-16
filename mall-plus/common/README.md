# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [cart/cart.proto](#cart/cart.proto)
    - [AddReq](#cart.AddReq)
    - [CartItem](#cart.CartItem)
    - [CartsReply](#cart.CartsReply)
    - [EditReq](#cart.EditReq)
    - [SkuReq](#cart.SkuReq)
    - [SkusReq](#cart.SkusReq)
  
    - [Cart](#cart.Cart)
  
- [center/user.proto](#center/user.proto)
    - [EditPwdReq](#center.EditPwdReq)
    - [EditReq](#center.EditReq)
    - [OfflineReq](#center.OfflineReq)
    - [OnlineReply](#center.OnlineReply)
    - [OnlineReq](#center.OnlineReq)
    - [PhoneReq](#center.PhoneReq)
    - [RegisterReply](#center.RegisterReply)
    - [RegisterReq](#center.RegisterReq)
    - [ServerIDReply](#center.ServerIDReply)
    - [ServerIDsReply](#center.ServerIDsReply)
    - [UIDReq](#center.UIDReq)
    - [UIDsReq](#center.UIDsReq)
    - [UserToken](#center.UserToken)
    - [Userinfo](#center.Userinfo)
    - [UsernameReq](#center.UsernameReq)
  
    - [Userinfo.Gender](#center.Userinfo.Gender)
  
    - [User](#center.User)
  
- [market/market.proto](#market/market.proto)
    - [AppSetting](#market.AppSetting)
    - [AppSettingReply](#market.AppSettingReply)
    - [CatReq](#market.CatReq)
    - [Coupon](#market.Coupon)
    - [CouponInfoReq](#market.CouponInfoReq)
    - [CouponInternal](#market.CouponInternal)
    - [CouponListReply](#market.CouponListReply)
    - [CouponReq](#market.CouponReq)
    - [CouponUsedReq](#market.CouponUsedReq)
    - [HomeCatDataReply](#market.HomeCatDataReply)
    - [HomeDataItem](#market.HomeDataItem)
    - [HomeDataReply](#market.HomeDataReply)
    - [Notice](#market.Notice)
    - [NoticeReply](#market.NoticeReply)
    - [PageReq](#market.PageReq)
    - [PayItem](#market.PayItem)
    - [PayReply](#market.PayReply)
    - [SearchReply](#market.SearchReply)
    - [SettingAds](#market.SettingAds)
    - [SettingImages](#market.SettingImages)
    - [SettingNav](#market.SettingNav)
    - [SettingNavs](#market.SettingNavs)
    - [SettingProduct](#market.SettingProduct)
    - [SkuReq](#market.SkuReq)
  
    - [Market](#market.Market)
  
- [member/member.proto](#member/member.proto)
    - [Address](#member.Address)
    - [AddressAddReq](#member.AddressAddReq)
    - [AddressIDReply](#member.AddressIDReply)
    - [AddressIDReq](#member.AddressIDReq)
    - [AddressInfoInternal](#member.AddressInfoInternal)
    - [AddressInfoReq](#member.AddressInfoReq)
    - [AddressReply](#member.AddressReply)
    - [CodeReply](#member.CodeReply)
    - [LoginReq](#member.LoginReq)
    - [MemberEditReq](#member.MemberEditReq)
    - [MemberInfo](#member.MemberInfo)
    - [MemberInfoReply](#member.MemberInfoReply)
    - [MemberToken](#member.MemberToken)
    - [MemberTokenReply](#member.MemberTokenReply)
    - [PhoneLoginReq](#member.PhoneLoginReq)
    - [PhoneReq](#member.PhoneReq)
    - [PwdEditReq](#member.PwdEditReq)
    - [RegisterReq](#member.RegisterReq)
  
    - [Member](#member.Member)
  
- [order/event.proto](#order/event.proto)
    - [Event](#order.Event)
  
- [order/order.proto](#order/order.proto)
    - [Address](#order.Address)
    - [CommentReq](#order.CommentReq)
    - [ListReply](#order.ListReply)
    - [ListReq](#order.ListReq)
    - [OrderIDReply](#order.OrderIDReply)
    - [OrderIDReq](#order.OrderIDReq)
    - [OrderInfo](#order.OrderInfo)
    - [OrderInfoReply](#order.OrderInfoReply)
    - [OrderList](#order.OrderList)
    - [OrderSku](#order.OrderSku)
    - [PayNotifyReq](#order.PayNotifyReq)
    - [RefundReq](#order.RefundReq)
    - [SkuSubmitReq](#order.SkuSubmitReq)
    - [SubmitReq](#order.SubmitReq)
  
    - [Order](#order.Order)
  
- [product/product.proto](#product/product.proto)
    - [Attr](#product.Attr)
    - [AttrEs](#product.AttrEs)
    - [Attrs](#product.Attrs)
    - [BrandEs](#product.BrandEs)
    - [CatEs](#product.CatEs)
    - [Category](#product.Category)
    - [CategoryReply](#product.CategoryReply)
    - [CommentReq](#product.CommentReq)
    - [SaleAttrs](#product.SaleAttrs)
    - [SearchAttrs](#product.SearchAttrs)
    - [SearchReply](#product.SearchReply)
    - [SearchReq](#product.SearchReq)
    - [Sku](#product.Sku)
    - [SkuAttr](#product.SkuAttr)
    - [SkuEs](#product.SkuEs)
    - [SkuInfo](#product.SkuInfo)
    - [SkuInfoInternal](#product.SkuInfoInternal)
    - [SkuListReply](#product.SkuListReply)
    - [SkuListReq](#product.SkuListReq)
    - [SkuReply](#product.SkuReply)
    - [SkuReq](#product.SkuReq)
    - [SkuSaleAttr](#product.SkuSaleAttr)
    - [SkuSaleAttrReply](#product.SkuSaleAttrReply)
    - [SkuValue](#product.SkuValue)
    - [Skus](#product.Skus)
  
    - [Product](#product.Product)
  
- [seckill/seckill.proto](#seckill/seckill.proto)
    - [KillReply](#seckill.KillReply)
    - [KillReq](#seckill.KillReq)
    - [Session](#seckill.Session)
    - [SessionIdReq](#seckill.SessionIdReq)
    - [SessionsReply](#seckill.SessionsReply)
    - [Sku](#seckill.Sku)
    - [SkuIdReq](#seckill.SkuIdReq)
    - [SkuReply](#seckill.SkuReply)
    - [SkusReply](#seckill.SkusReply)
  
    - [Seckill](#seckill.Seckill)
  
- [task/msg.proto](#task/msg.proto)
    - [Msg](#task.Msg)
  
    - [Msg.Type](#task.Msg.Type)
  
- [third/third.proto](#third/third.proto)
    - [CodeReply](#third.CodeReply)
    - [ETHPayReq](#third.ETHPayReq)
    - [PhoneReq](#third.PhoneReq)
    - [VCodeReq](#third.VCodeReq)
  
    - [Third](#third.Third)
  
- [warehouse/event.proto](#warehouse/event.proto)
    - [Event](#warehouse.Event)
  
- [warehouse/warehouse.proto](#warehouse/warehouse.proto)
    - [SkuStockLockReq](#warehouse.SkuStockLockReq)
    - [SkuStockLockReq.SkuNumEntry](#warehouse.SkuStockLockReq.SkuNumEntry)
    - [SkuStockNum](#warehouse.SkuStockNum)
    - [SkuStockNum.SkuNumEntry](#warehouse.SkuStockNum.SkuNumEntry)
    - [SkuStockReq](#warehouse.SkuStockReq)
    - [SkuStockUnlockReq](#warehouse.SkuStockUnlockReq)
    - [SpuStockReq](#warehouse.SpuStockReq)
    - [StockNumReply](#warehouse.StockNumReply)
  
    - [Warehouse](#warehouse.Warehouse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cart/cart.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cart/cart.proto
购物车服务


<a name="cart.AddReq"></a>

### AddReq
添加购物车请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |
| num | [int32](#int32) |  | 数量 |






<a name="cart.CartItem"></a>

### CartItem
购物车结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |
| title | [string](#string) |  | 商品标题 |
| price | [double](#double) |  | 商品价格 |
| cover | [string](#string) |  | 商品封面 |
| sku_attr | [string](#string) |  | 商品销售属性 |
| num | [int32](#int32) |  | 数量 |






<a name="cart.CartsReply"></a>

### CartsReply
购物车列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [CartItem](#cart.CartItem) | repeated |  |






<a name="cart.EditReq"></a>

### EditReq
修改购物车购微项请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| old_sku_id | [int64](#int64) |  | 修改前商品id |
| new_sku_id | [int64](#int64) |  | 修改后商品id |
| num | [int32](#int32) |  | 数量 |






<a name="cart.SkuReq"></a>

### SkuReq
sku_id请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |






<a name="cart.SkusReq"></a>

### SkusReq
多sku_id请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| sku_ids | [int64](#int64) | repeated | sku_id数组 |





 

 

 


<a name="cart.Cart"></a>

### Cart
购物车服务接口给定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddCart | [AddReq](#cart.AddReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 添加购物车 |
| EditCart | [EditReq](#cart.EditReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 更新购物车 |
| EditCartNum | [AddReq](#cart.AddReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 更新购物车数量 |
| DelCart | [SkuReq](#cart.SkuReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 删除购物项 |
| ClearCart | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | 清空购物车 |
| MyCart | [.google.protobuf.Empty](#google.protobuf.Empty) | [CartsReply](#cart.CartsReply) | 我的购物车 |
| BatchGetCarts | [SkusReq](#cart.SkusReq) | [CartsReply](#cart.CartsReply) | 批量获取购物车信息 |
| BatchDelCart | [SkusReq](#cart.SkusReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 批量删除购物车 |

 



<a name="center/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## center/user.proto
中心服务


<a name="center.EditPwdReq"></a>

### EditPwdReq
用户密码修改请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |
| old_pwd | [string](#string) |  | 原密码 |
| pwd | [string](#string) |  | 新密码 |






<a name="center.EditReq"></a>

### EditReq
用户信息修改请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |
| content | [bytes](#bytes) |  | json信息体 |






<a name="center.OfflineReq"></a>

### OfflineReq
用户下线请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | 用户id |
| key | [string](#string) |  | 键 |
| server | [string](#string) |  | 服务器id |






<a name="center.OnlineReply"></a>

### OnlineReply
用户上线响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | 用户id |
| key | [string](#string) |  | 键 |






<a name="center.OnlineReq"></a>

### OnlineReq
用户上线请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [string](#string) |  | 服务器id |
| token | [string](#string) |  | 用户令牌 |






<a name="center.PhoneReq"></a>

### PhoneReq
手机号登录请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [int64](#int64) |  | 手机号 |






<a name="center.RegisterReply"></a>

### RegisterReply
注册响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |






<a name="center.RegisterReq"></a>

### RegisterReq
注册请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| phone | [int64](#int64) |  | 手机号 |






<a name="center.ServerIDReply"></a>

### ServerIDReply
获取用户所有服务器id响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serverID | [string](#string) |  |  |






<a name="center.ServerIDsReply"></a>

### ServerIDsReply
批量获取用户所有服务器id响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serverIDs | [string](#string) | repeated |  |






<a name="center.UIDReq"></a>

### UIDReq
用户id请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |






<a name="center.UIDsReq"></a>

### UIDsReq
用户id数组结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int64](#int64) | repeated | 用户id数组 |






<a name="center.UserToken"></a>

### UserToken
用户令牌信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [Userinfo](#center.Userinfo) |  |  |
| token | [string](#string) |  | 令牌 |






<a name="center.Userinfo"></a>

### Userinfo
用户基础信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |
| username | [string](#string) |  | 用户名 |
| nickname | [string](#string) |  | 昵称 |
| phone | [int64](#int64) |  | 手机号 |
| email | [string](#string) |  | 邮箱 |
| sign | [string](#string) |  | 用户签名 |
| avatar | [string](#string) |  | 用户头像 |
| area | [string](#string) |  | 地区信息 |
| gender | [Userinfo.Gender](#center.Userinfo.Gender) |  | 性别 |






<a name="center.UsernameReq"></a>

### UsernameReq
用户名登录请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |





 


<a name="center.Userinfo.Gender"></a>

### Userinfo.Gender


| Name | Number | Description |
| ---- | ------ | ----------- |
| MALE | 0 | 男 |
| FEMALE | 1 | 女 |
| SECRET | 2 | 保密 |


 

 


<a name="center.User"></a>

### User
用户中心服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterReq](#center.RegisterReq) | [RegisterReply](#center.RegisterReply) | 用户注册 |
| UsernameLogin | [UsernameReq](#center.UsernameReq) | [UserToken](#center.UserToken) | 用户名密码登录 |
| PhoneLogin | [PhoneReq](#center.PhoneReq) | [UserToken](#center.UserToken) | 手机号登录 |
| Edit | [EditReq](#center.EditReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 修改用户信息 |
| EditPwd | [EditPwdReq](#center.EditPwdReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 修改密码 |
| Info | [UIDReq](#center.UIDReq) | [Userinfo](#center.Userinfo) | 获取用户信息 |
| Logout | [UIDReq](#center.UIDReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 用户登出 |
| Online | [OnlineReq](#center.OnlineReq) | [OnlineReply](#center.OnlineReply) | 用户上线，建立长连接 |
| Offline | [OfflineReq](#center.OfflineReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 用户下线，断开长连接 |
| ServerID | [UIDReq](#center.UIDReq) | [ServerIDReply](#center.ServerIDReply) | 获取用户长连接所在的服务器ID |
| BatchServersIDs | [UIDsReq](#center.UIDsReq) | [ServerIDsReply](#center.ServerIDsReply) | 批量获取长连接所在的服务器ID |

 



<a name="market/market.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## market/market.proto
营销服务


<a name="market.AppSetting"></a>

### AppSetting
页面配置数据结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int32](#int32) |  | 配置类型 |
| images | [SettingImages](#market.SettingImages) |  |  |
| navs | [SettingNavs](#market.SettingNavs) |  |  |
| ads | [SettingAds](#market.SettingAds) |  |  |
| product | [SettingProduct](#market.SettingProduct) |  |  |






<a name="market.AppSettingReply"></a>

### AppSettingReply
页面配置列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [AppSetting](#market.AppSetting) | repeated |  |






<a name="market.CatReq"></a>

### CatReq
分类id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cat_id | [int64](#int64) |  |  |






<a name="market.Coupon"></a>

### Coupon
优惠券结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 优惠券id |
| name | [string](#string) |  | 优惠券名称 |
| amount | [double](#double) |  | 优惠券面值 |
| min_point | [double](#double) |  | 使用门槛 |
| start_at | [int64](#int64) |  | 有效开始时间 |
| end_at | [int64](#int64) |  | 有效结束时间 |
| note | [string](#string) |  | 描述 |
| status | [int32](#int32) |  | 状态 |






<a name="market.CouponInfoReq"></a>

### CouponInfoReq
优惠券详情请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户id |
| coupon_id | [int64](#int64) |  | 会员优惠券id |






<a name="market.CouponInternal"></a>

### CouponInternal
---- 内部响应 ----
优惠券详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [Coupon](#market.Coupon) |  |  |






<a name="market.CouponListReply"></a>

### CouponListReply
优惠券列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Coupon](#market.Coupon) | repeated |  |






<a name="market.CouponReq"></a>

### CouponReq
优惠券请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coupon_id | [int64](#int64) |  | 优惠券id |






<a name="market.CouponUsedReq"></a>

### CouponUsedReq
优惠券使用请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户id |
| coupon_id | [int64](#int64) |  | 优惠券id |
| order_id | [int64](#int64) |  | 订单id |






<a name="market.HomeCatDataReply"></a>

### HomeCatDataReply
首页分类数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [AppSetting](#market.AppSetting) | repeated | 分类下配置页面数据 |






<a name="market.HomeDataItem"></a>

### HomeDataItem
首页数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 分类id |
| name | [string](#string) |  | 分类名 |
| list | [AppSetting](#market.AppSetting) | repeated | 分类下配置页面数据 |






<a name="market.HomeDataReply"></a>

### HomeDataReply
首页数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [HomeDataItem](#market.HomeDataItem) | repeated |  |






<a name="market.Notice"></a>

### Notice
公告结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 公告id |
| title | [string](#string) |  | 公告标题 |
| content | [string](#string) |  | 公告内容 |
| created_at | [int64](#int64) |  | 公告创建时间 |






<a name="market.NoticeReply"></a>

### NoticeReply
公告列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Notice](#market.Notice) | repeated | 公告列表 |






<a name="market.PageReq"></a>

### PageReq
分页结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 第几页 |






<a name="market.PayItem"></a>

### PayItem
支付配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 支付id |
| name | [string](#string) |  | 支付名称 |
| address | [string](#string) |  | 支付地址 |






<a name="market.PayReply"></a>

### PayReply
支付配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PayItem](#market.PayItem) | repeated |  |






<a name="market.SearchReply"></a>

### SearchReply
搜索页配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [AppSetting](#market.AppSetting) | repeated | 搜索页配置数据 |
| hot | [string](#string) | repeated | 搜索热词 |






<a name="market.SettingAds"></a>

### SettingAds
单图广告结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 标题 |
| cover | [string](#string) |  | 广告图 |






<a name="market.SettingImages"></a>

### SettingImages
页面配置类型
图片组结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [string](#string) | repeated |  |






<a name="market.SettingNav"></a>

### SettingNav
图标结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 标题 |
| icon | [string](#string) |  | icon |






<a name="market.SettingNavs"></a>

### SettingNavs
图标组


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [SettingNav](#market.SettingNav) | repeated |  |






<a name="market.SettingProduct"></a>

### SettingProduct
产品列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| router | [string](#string) |  | 路由 |






<a name="market.SkuReq"></a>

### SkuReq
商品请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |





 

 

 


<a name="market.Market"></a>

### Market
营销服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetHomeData | [.google.protobuf.Empty](#google.protobuf.Empty) | [HomeDataReply](#market.HomeDataReply) | 获取首页配置数据 |
| GetHomeCatData | [CatReq](#market.CatReq) | [AppSettingReply](#market.AppSettingReply) | 获取首页分类下配置数据 |
| GetNoticeList | [PageReq](#market.PageReq) | [NoticeReply](#market.NoticeReply) | 获取公告列表 |
| GetSearchData | [.google.protobuf.Empty](#google.protobuf.Empty) | [SearchReply](#market.SearchReply) | 获取搜索页配置数据 |
| GetPayConfig | [.google.protobuf.Empty](#google.protobuf.Empty) | [PayReply](#market.PayReply) | 获取支付配置 |
| GetCouponList | [SkuReq](#market.SkuReq) | [CouponListReply](#market.CouponListReply) | 商品可以领取的优惠券列表 |
| GetMyCouponList | [.google.protobuf.Empty](#google.protobuf.Empty) | [CouponListReply](#market.CouponListReply) | 我的优惠券列表 |
| CouponDraw | [CouponReq](#market.CouponReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 领取优惠券 |
| CouponUsed | [CouponUsedReq](#market.CouponUsedReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | ---- 以下内部调用 ---- / 使用优惠券 |
| GetCouponInfo | [CouponInfoReq](#market.CouponInfoReq) | [CouponInternal](#market.CouponInternal) | 获取优惠券详情 |

 



<a name="member/member.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## member/member.proto
会员服务


<a name="member.Address"></a>

### Address
收货地址结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 主键 |
| name | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| province | [string](#string) |  | 省 |
| city | [string](#string) |  | 市 |
| county | [string](#string) |  | 区/县 |
| detail | [string](#string) |  | 详细地址 |
| area_code | [int64](#int64) |  | 地区码 |
| is_default | [int32](#int32) |  | 是否设置默认地址 |






<a name="member.AddressAddReq"></a>

### AddressAddReq
添加地址结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| province | [string](#string) |  | 省 |
| city | [string](#string) |  | 市 |
| county | [string](#string) |  | 区/县 |
| detail | [string](#string) |  | 详细地址 |
| area_code | [int64](#int64) |  | 地区码 |
| is_default | [int32](#int32) |  | 是否设置默认地址 |






<a name="member.AddressIDReply"></a>

### AddressIDReply
收货地址id结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [int64](#int64) |  | 收货地址id |






<a name="member.AddressIDReq"></a>

### AddressIDReq
收货地址id结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 收货地址id |






<a name="member.AddressInfoInternal"></a>

### AddressInfoInternal
收货地址详情详情结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [Address](#member.Address) |  |  |






<a name="member.AddressInfoReq"></a>

### AddressInfoReq
获取收货地址详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户id |
| address_id | [int64](#int64) |  | 收货地址id |






<a name="member.AddressReply"></a>

### AddressReply
收货地址列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Address](#member.Address) | repeated |  |






<a name="member.CodeReply"></a>

### CodeReply
发送短信响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [string](#string) |  | 验证码 |






<a name="member.LoginReq"></a>

### LoginReq
登录请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |






<a name="member.MemberEditReq"></a>

### MemberEditReq
会员信息修改请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| sign | [string](#string) |  |  |






<a name="member.MemberInfo"></a>

### MemberInfo
会员基础信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户id |
| username | [string](#string) |  | 用户名 |
| nickname | [string](#string) |  | 昵称 |
| sign | [string](#string) |  | 用户签名 |
| avatar | [string](#string) |  | 用户头像 |
| area | [string](#string) |  | 地区信息 |
| phone | [string](#string) |  | 手机号 |






<a name="member.MemberInfoReply"></a>

### MemberInfoReply
会员信息响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [MemberInfo](#member.MemberInfo) |  |  |






<a name="member.MemberToken"></a>

### MemberToken
用户令牌信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member | [MemberInfo](#member.MemberInfo) |  | 会员信息 |
| token | [string](#string) |  | 令牌 |






<a name="member.MemberTokenReply"></a>

### MemberTokenReply
登录成功令牌会员信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [MemberToken](#member.MemberToken) |  |  |






<a name="member.PhoneLoginReq"></a>

### PhoneLoginReq
手机号登录请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 手机号 |
| code | [string](#string) |  | 验证码 |






<a name="member.PhoneReq"></a>

### PhoneReq
手机号请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 手机号 |






<a name="member.PwdEditReq"></a>

### PwdEditReq
修改密码请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| old_password | [string](#string) |  | 原密码 |
| password | [string](#string) |  | 密码 |
| confirm_password | [string](#string) |  | 确认密码 |






<a name="member.RegisterReq"></a>

### RegisterReq
注册请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| phone | [string](#string) |  | 手机号 |
| code | [string](#string) |  | 验证码 |
| password | [string](#string) |  | 密码 |
| confirm_password | [string](#string) |  | 确认密码 |





 

 

 


<a name="member.Member"></a>

### Member
会员服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterReq](#member.RegisterReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 注册 |
| Login | [LoginReq](#member.LoginReq) | [MemberTokenReply](#member.MemberTokenReply) | 用户名密码登录 |
| PhoneLogin | [PhoneLoginReq](#member.PhoneLoginReq) | [MemberTokenReply](#member.MemberTokenReply) | 手机号登录 |
| MemberEdit | [MemberEditReq](#member.MemberEditReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 修改会员信息 |
| MemberPwdEdit | [PwdEditReq](#member.PwdEditReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 修改密码 |
| MemberProfile | [.google.protobuf.Empty](#google.protobuf.Empty) | [MemberInfoReply](#member.MemberInfoReply) | 获取会员信息 |
| Logout | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | 登出 |
| AddressAdd | [AddressAddReq](#member.AddressAddReq) | [AddressIDReply](#member.AddressIDReply) | 添加收货地址 |
| AddressEdit | [Address](#member.Address) | [.google.protobuf.Empty](#google.protobuf.Empty) | 修改收货地址 |
| GetAddressList | [.google.protobuf.Empty](#google.protobuf.Empty) | [AddressReply](#member.AddressReply) | 收货地址列表 |
| AddressDel | [AddressIDReq](#member.AddressIDReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 删除收货地址 |
| SendCode | [PhoneReq](#member.PhoneReq) | [CodeReply](#member.CodeReply) | 发送短信验证码 |
| GetAddressInfo | [AddressInfoReq](#member.AddressInfoReq) | [AddressInfoInternal](#member.AddressInfoInternal) | ---- 以下内部调用 ---- / 获取收货地址信息 |

 



<a name="order/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order/event.proto



<a name="order.Event"></a>

### Event
订单服务秒杀消息结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [int64](#int64) |  | 会员id |
| sku_id | [int64](#int64) |  | 商品id |
| address_id | [int64](#int64) |  | 收货地址id |
| num | [int32](#int32) |  | 秒杀数量 |
| price | [int32](#int32) |  | 秒杀价格 |
| order_no | [string](#string) |  | 订单号 |





 

 

 

 



<a name="order/order.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order/order.proto
订单服务


<a name="order.Address"></a>

### Address
收货地址信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| area | [string](#string) |  | 地区 |
| detail | [string](#string) |  | 详细地址 |






<a name="order.CommentReq"></a>

### CommentReq
评价请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| star | [int32](#int32) |  | 评分 |
| order_id | [int64](#int64) |  | 订单id |
| content | [string](#string) |  | 评价内容 |
| resources | [string](#string) |  | 评论资源 |
| sku_ids | [int64](#int64) | repeated | 评价商品 |






<a name="order.ListReply"></a>

### ListReply
订单列表响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [OrderList](#order.OrderList) | repeated |  |






<a name="order.ListReq"></a>

### ListReq
订单列表请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [int32](#int32) |  | 订单状态 |
| page | [int32](#int32) |  | 订单页码 |






<a name="order.OrderIDReply"></a>

### OrderIDReply
订单id结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [int64](#int64) |  | 订单id |






<a name="order.OrderIDReq"></a>

### OrderIDReq
订单id结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |






<a name="order.OrderInfo"></a>

### OrderInfo
订单详情结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 订单id |
| order_no | [string](#string) |  | 订单号 |
| note | [string](#string) |  | 订单会员备注 |
| total_amount | [double](#double) |  | 订单总价 |
| amount | [double](#double) |  | 订单金额 |
| coupon_amount | [double](#double) |  | 优惠券优惠金额 |
| freight_amount | [double](#double) |  | 运费 |
| pay_amount | [double](#double) |  | 支付金额 |
| pay_type | [int32](#int32) |  | 支付类型 |
| pay_at | [int64](#int64) |  | 支付时间 |
| create_at | [int64](#int64) |  | 订单创建时间 |
| status | [int32](#int32) |  | 订单状态 |
| trade_no | [string](#string) |  | 支付交易流水号 |
| delivery_company | [string](#string) |  | 物流公司 |
| delivery_no | [string](#string) |  | 物流单号 |
| integration | [int32](#int32) |  | 所获积分 |
| growth | [int32](#int32) |  | 所获得成长值 |
| delivery_at | [int64](#int64) |  | 发货时间 |
| receive_at | [int64](#int64) |  | 确认收货时间 |
| comment_at | [int64](#int64) |  | 评价时间 |
| address | [Address](#order.Address) |  | 收货地址信息 |
| items | [OrderSku](#order.OrderSku) | repeated | 订单对应的商品列表 |






<a name="order.OrderInfoReply"></a>

### OrderInfoReply
订单详情响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [OrderInfo](#order.OrderInfo) |  |  |






<a name="order.OrderList"></a>

### OrderList
订单列表结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 订单id |
| order_no | [string](#string) |  | 订单号 |
| amount | [double](#double) |  | 订单金额 |
| status | [int32](#int32) |  | 订单状态 |
| time | [int64](#int64) |  | 创建时间 |
| items | [OrderSku](#order.OrderSku) | repeated | 订单对应的商品列表 |






<a name="order.OrderSku"></a>

### OrderSku
订单商品信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |
| title | [string](#string) |  | 商品标题 |
| cover | [string](#string) |  | 封面图 |
| price | [double](#double) |  | 价格 |
| num | [int32](#int32) |  | 数量 |
| attr_value | [string](#string) |  | 销售属性值 |






<a name="order.PayNotifyReq"></a>

### PayNotifyReq
支付回调请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pay_amount | [int64](#int64) |  | 支付金额 |
| pay_type | [int64](#int64) |  | 支付类型 |
| order_no | [string](#string) |  | 订单号 |
| trade_no | [string](#string) |  | 交易号 |
| trans_hash | [string](#string) |  | 交易hash |






<a name="order.RefundReq"></a>

### RefundReq
申请退款请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| content | [string](#string) |  | 理由 |






<a name="order.SkuSubmitReq"></a>

### SkuSubmitReq
商品直接提交订单请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |
| address_id | [int64](#int64) |  | 收货地址id |
| coupon_id | [int64](#int64) |  | 优惠券id |
| note | [string](#string) |  | 用户备注 |
| num | [int32](#int32) |  | 购买数量 |






<a name="order.SubmitReq"></a>

### SubmitReq
购物车提交订单请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_ids | [int64](#int64) | repeated | 购物车sku_id列表 |
| address_id | [int64](#int64) |  | 收货地址id |
| coupon_id | [int64](#int64) |  | 优惠券id |
| note | [string](#string) |  | 用户备注 |





 

 

 


<a name="order.Order"></a>

### Order
订单服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SubmitOrder | [SubmitReq](#order.SubmitReq) | [OrderIDReply](#order.OrderIDReply) | 从购物车提交订单 |
| SubmitSkuOrder | [SkuSubmitReq](#order.SkuSubmitReq) | [OrderIDReply](#order.OrderIDReply) | 商品直接提交订单 |
| OrderDetail | [OrderIDReq](#order.OrderIDReq) | [OrderInfoReply](#order.OrderInfoReply) | 订单详情 |
| OrderCancel | [OrderIDReq](#order.OrderIDReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 订单取消 |
| OrderList | [ListReq](#order.ListReq) | [ListReply](#order.ListReply) | 订单列表 |
| OrderPayNotify | [PayNotifyReq](#order.PayNotifyReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 订单支付成功回调 |
| OrderRefund | [RefundReq](#order.RefundReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 订单退款 |
| OrderConfirmReceipt | [OrderIDReq](#order.OrderIDReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 订单确认收货 |
| OrderComment | [CommentReq](#order.CommentReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 订单评价 |
| GetOrderByID | [OrderIDReq](#order.OrderIDReq) | [OrderInfo](#order.OrderInfo) | 内部调用 订单信息 |

 



<a name="product/product.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product/product.proto
产品服务


<a name="product.Attr"></a>

### Attr
属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 属性id |
| name | [string](#string) |  | 属性名 |
| value | [string](#string) |  | 属性值 |






<a name="product.AttrEs"></a>

### AttrEs
es中属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 属性id |
| name | [string](#string) |  | 属性名 |
| values | [string](#string) | repeated | 属性值 |






<a name="product.Attrs"></a>

### Attrs
属性分组以及分组下规格属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int64](#int64) |  | 规格属性分组id |
| group_name | [string](#string) |  | 规格属性分组名 |
| items | [Attr](#product.Attr) | repeated | 分组下所有规格属性 |






<a name="product.BrandEs"></a>

### BrandEs
es中品牌结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 品牌id |
| name | [string](#string) |  | 品牌名 |
| logo | [string](#string) |  | 品牌logo |






<a name="product.CatEs"></a>

### CatEs
es中分类结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 分类id |
| name | [string](#string) |  | 分类名 |






<a name="product.Category"></a>

### Category
产品分类结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 分类id |
| parent_id | [int64](#int64) |  | 父id |
| name | [string](#string) |  | 分类名 |
| sort | [int32](#int32) |  | 排序值 |
| child | [Category](#product.Category) | repeated | 下级分类列表 |






<a name="product.CategoryReply"></a>

### CategoryReply
产品分类


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Category](#product.Category) | repeated |  |






<a name="product.CommentReq"></a>

### CommentReq
---- 内部请求 ----
评论请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_ids | [int64](#int64) | repeated | sku_id列表 |
| user_id | [int64](#int64) |  | 用户id |
| order_id | [int64](#int64) |  | 订单id |
| star | [int32](#int32) |  | 星级 |
| content | [string](#string) |  | 评价内容 |
| resources | [string](#string) |  | 评价资源 |






<a name="product.SaleAttrs"></a>

### SaleAttrs
sku下所有销售属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attr_id | [int64](#int64) |  | 销售属性id |
| attr_name | [string](#string) |  | 销售属性名 |
| values | [SkuValue](#product.SkuValue) | repeated | 属性值列表 |






<a name="product.SearchAttrs"></a>

### SearchAttrs
搜索规格结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 规格名id |
| values | [string](#string) | repeated | 规格值列表 |






<a name="product.SearchReply"></a>

### SearchReply
搜索结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [SkuEs](#product.SkuEs) | repeated | 搜索商品 |
| brands | [BrandEs](#product.BrandEs) | repeated | 当前查询到的结果锁涉及到的品牌 |
| attrs | [AttrEs](#product.AttrEs) | repeated | 当前查询到的结果锁涉及到的所有属性 |
| cats | [CatEs](#product.CatEs) | repeated | 当前查询到的结果锁涉及到的所有分类 |






<a name="product.SearchReq"></a>

### SearchReq
搜索请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyword | [string](#string) |  | 关键字 |
| cat_id | [int64](#int64) |  | 分类id |
| field | [int32](#int32) |  | 排序字段 |
| order | [int32](#int32) |  | 排序类型 0=asc 1=desc |
| has_stock | [bool](#bool) |  | 是否有库存 |
| price_s | [int32](#int32) |  | 价格区间起始 |
| price_e | [int32](#int32) |  | 价格区间止 |
| brand_id | [int64](#int64) | repeated | 品牌,多选 |
| attrs | [SearchAttrs](#product.SearchAttrs) | repeated | 属性 eg: 1_桌子,椅子 |
| page | [int32](#int32) |  | 分页 |






<a name="product.Sku"></a>

### Sku
商品详情结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | sku_id |
| spu_id | [int64](#int64) |  | spu_id |
| cat_id | [int64](#int64) |  | 分类id |
| brand_id | [int64](#int64) |  | 品牌id |
| title | [string](#string) |  | 商品标题 |
| desc | [string](#string) |  | 描述 |
| cover | [string](#string) |  | 封面图 |
| subtitle | [string](#string) |  | 副标题 |
| price | [double](#double) |  | 价格 |
| sale_count | [int64](#int64) |  | 销量 |
| stock | [int32](#int32) |  | 库存 |
| is_many | [bool](#bool) |  | 是否多规格 |
| skus | [Skus](#product.Skus) | repeated | spu下所有sku商品 |
| attrs | [Attrs](#product.Attrs) | repeated | 当前sku对应spu规格属性 |
| sale_attrs | [SaleAttrs](#product.SaleAttrs) | repeated | 当前sku下的销售属性 |
| banners | [string](#string) | repeated | sku图集 |
| mains | [string](#string) | repeated | spu介绍 |






<a name="product.SkuAttr"></a>

### SkuAttr
销售属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attr_id | [int64](#int64) |  | 属性id |
| value_id | [int64](#int64) |  | 属性值id |
| attr_name | [string](#string) |  | 属性名 |
| value_name | [string](#string) |  | 属性值 |






<a name="product.SkuEs"></a>

### SkuEs
es中sku结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | sku_id |
| title | [string](#string) |  | 商品标题 |
| price | [double](#double) |  | 价格 |
| cover | [string](#string) |  | 封面 |
| sale_count | [int32](#int32) |  | 销量 |
| has_stock | [bool](#bool) |  | 是否有库存 |






<a name="product.SkuInfo"></a>

### SkuInfo
sku商品信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | sku_id |
| spu_id | [int64](#int64) |  | spu_id |
| cat_id | [int64](#int64) |  | 分类id |
| brand_id | [int64](#int64) |  | 品牌id |
| title | [string](#string) |  | 商品标题 |
| desc | [string](#string) |  | 描述 |
| cover | [string](#string) |  | 封面图 |
| subtitle | [string](#string) |  | 副标题 |
| price | [int64](#int64) |  | 价格 |
| sale_count | [int64](#int64) |  | 销量 |
| attr_value | [string](#string) |  | 销售属性值 |






<a name="product.SkuInfoInternal"></a>

### SkuInfoInternal
---- 内部响应 ----
sku商品信息响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [SkuInfo](#product.SkuInfo) |  |  |






<a name="product.SkuListReply"></a>

### SkuListReply
产品列表响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [SkuEs](#product.SkuEs) | repeated | 产品列表 |






<a name="product.SkuListReq"></a>

### SkuListReq
商品列表请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cat_id | [int64](#int64) |  | 分类 |
| page | [int32](#int32) |  | 分页 |






<a name="product.SkuReply"></a>

### SkuReply
商品详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Sku](#product.Sku) |  |  |






<a name="product.SkuReq"></a>

### SkuReq
商品详情请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  |  |






<a name="product.SkuSaleAttr"></a>

### SkuSaleAttr
sku销售属性结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | sku_id |
| is_many | [bool](#bool) |  | 是否多规格 |
| skus | [Skus](#product.Skus) | repeated | spu下所有sku商品 |
| sale_attrs | [SaleAttrs](#product.SaleAttrs) | repeated | 当前sku下的销售属性 |






<a name="product.SkuSaleAttrReply"></a>

### SkuSaleAttrReply
sku销售属性


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [SkuSaleAttr](#product.SkuSaleAttr) |  |  |






<a name="product.SkuValue"></a>

### SkuValue
规格值结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 属性值id |
| name | [string](#string) |  | 属性名 |






<a name="product.Skus"></a>

### Skus
spu下所有sku


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |
| price | [double](#double) |  | 价格 |
| stock | [int32](#int32) |  | 库存 |
| attrs | [SkuAttr](#product.SkuAttr) | repeated | 对应的销售属性 |





 

 

 


<a name="product.Product"></a>

### Product
产品服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CategoryTree | [.google.protobuf.Empty](#google.protobuf.Empty) | [CategoryReply](#product.CategoryReply) | 获取产品三级分类树 |
| SkuSearch | [SearchReq](#product.SearchReq) | [SearchReply](#product.SearchReply) | sku商品搜索 |
| SkuList | [SkuListReq](#product.SkuListReq) | [SkuListReply](#product.SkuListReply) | sku商品列表 |
| SkuDetail | [SkuReq](#product.SkuReq) | [SkuReply](#product.SkuReply) | sku商品详情 |
| GetSkuSaleAttrs | [SkuReq](#product.SkuReq) | [SkuSaleAttrReply](#product.SkuSaleAttrReply) | sku销售属性 |
| GetSkuByID | [SkuReq](#product.SkuReq) | [SkuInfoInternal](#product.SkuInfoInternal) | ---- 以下内部调用 ---- / sku信息 |
| SpuComment | [CommentReq](#product.CommentReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 商品评价 |

 



<a name="seckill/seckill.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## seckill/seckill.proto
秒杀服务


<a name="seckill.KillReply"></a>

### KillReply
秒杀响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [string](#string) |  | 订单号 |






<a name="seckill.KillReq"></a>

### KillReq
秒杀请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  |  |
| address_id | [int64](#int64) |  |  |
| num | [int64](#int64) |  |  |
| key | [string](#string) |  |  |






<a name="seckill.Session"></a>

### Session
场次信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 场次id |
| name | [string](#string) |  | 场次名 |
| open | [bool](#bool) |  | 是否正在秒杀 |
| skus | [Sku](#seckill.Sku) | repeated | 所有秒杀商品 |






<a name="seckill.SessionIdReq"></a>

### SessionIdReq
场次id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  | 场次id |






<a name="seckill.SessionsReply"></a>

### SessionsReply
秒杀场次响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Session](#seckill.Session) | repeated |  |






<a name="seckill.Sku"></a>

### Sku
秒杀商品信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 商品id |
| price | [double](#double) |  | 秒杀价格 |
| count | [int32](#int32) |  | 秒杀数量 |
| limit | [int32](#int32) |  | 个人限购 |
| original_price | [double](#double) |  | 原价 |
| title | [string](#string) |  | 标题 |
| cover | [string](#string) |  | 封面 |
| key | [string](#string) |  | 加密key |
| open | [bool](#bool) |  | 是否正在秒杀 |
| start_at | [int64](#int64) |  | 秒杀开始时间 |






<a name="seckill.SkuIdReq"></a>

### SkuIdReq
商品id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | 商品id |






<a name="seckill.SkuReply"></a>

### SkuReply
秒杀商品信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Sku](#seckill.Sku) |  |  |






<a name="seckill.SkusReply"></a>

### SkusReply
秒杀商品列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Sku](#seckill.Sku) | repeated |  |





 

 

 


<a name="seckill.Seckill"></a>

### Seckill
秒杀服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Kill | [KillReq](#seckill.KillReq) | [KillReply](#seckill.KillReply) | 秒杀 |
| GetSessionAll | [.google.protobuf.Empty](#google.protobuf.Empty) | [SessionsReply](#seckill.SessionsReply) | 获取所有秒杀场次 |
| GetSkusList | [SessionIdReq](#seckill.SessionIdReq) | [SkusReply](#seckill.SkusReply) | 获取场次下所有秒杀商品 |
| GetSkuByID | [SkuIdReq](#seckill.SkuIdReq) | [SkuReply](#seckill.SkuReply) | 获取商品秒杀详情 |

 



<a name="task/msg.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task/msg.proto



<a name="task.Msg"></a>

### Msg
消息格式


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Msg.Type](#task.Msg.Type) |  | 发送类型 |
| server | [string](#string) |  | 用户所在服务器id |
| event | [string](#string) |  | 消息类型 |
| userIds | [int64](#int64) | repeated |  |
| timestamp | [int64](#int64) |  |  |
| body | [bytes](#bytes) |  |  |





 


<a name="task.Msg.Type"></a>

### Msg.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEND | 0 | 指定用户消息 |
| BROADCAST | 1 | 广播消息 |
| CLOSE | 2 | 主动关闭客户端连接 |
| History | 3 | 历史消息 |


 

 

 



<a name="third/third.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## third/third.proto
第三方服务


<a name="third.CodeReply"></a>

### CodeReply
发送短信响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | 验证码 |






<a name="third.ETHPayReq"></a>

### ETHPayReq
以太币支付检测请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 支付合约id |
| address | [string](#string) |  | 支付合约地址 |
| orderNo | [string](#string) |  | 订单号 |






<a name="third.PhoneReq"></a>

### PhoneReq
手机号请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | 手机号 |






<a name="third.VCodeReq"></a>

### VCodeReq
检测手机验证码合法请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [int64](#int64) |  | 手机号 |
| code | [string](#string) |  | 验证码 |





 

 

 


<a name="third.Third"></a>

### Third
第三方服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendSMS | [PhoneReq](#third.PhoneReq) | [CodeReply](#third.CodeReply) | 发送短信验证码 |
| CheckVCode | [VCodeReq](#third.VCodeReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 短信验证码验证 |
| CheckETHPay | [ETHPayReq](#third.ETHPayReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 以太币支付检测 |

 



<a name="warehouse/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## warehouse/event.proto



<a name="warehouse.Event"></a>

### Event
仓储服务消息结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| finish | [bool](#bool) |  | 订单是否已完成 |





 

 

 

 



<a name="warehouse/warehouse.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## warehouse/warehouse.proto
仓储服务


<a name="warehouse.SkuStockLockReq"></a>

### SkuStockLockReq
锁定库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  |  |
| order_no | [string](#string) |  | 订单号 |
| consignee | [string](#string) |  | 收货人 |
| phone | [string](#string) |  | 收货人手机号 |
| address | [string](#string) |  | 收货地址 |
| note | [string](#string) |  | 订单备注 |
| sku_num | [SkuStockLockReq.SkuNumEntry](#warehouse.SkuStockLockReq.SkuNumEntry) | repeated | sku_id =&gt; 库存数量 |






<a name="warehouse.SkuStockLockReq.SkuNumEntry"></a>

### SkuStockLockReq.SkuNumEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int32](#int32) |  |  |






<a name="warehouse.SkuStockNum"></a>

### SkuStockNum
sku库存数量map结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_num | [SkuStockNum.SkuNumEntry](#warehouse.SkuStockNum.SkuNumEntry) | repeated | sku_id =&gt; 库存数量 |






<a name="warehouse.SkuStockNum.SkuNumEntry"></a>

### SkuStockNum.SkuNumEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int32](#int32) |  |  |






<a name="warehouse.SkuStockReq"></a>

### SkuStockReq
sku库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku_id | [int64](#int64) |  | sku_id |






<a name="warehouse.SkuStockUnlockReq"></a>

### SkuStockUnlockReq
解锁库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [int64](#int64) |  | 订单id |
| finish | [bool](#bool) |  | 订单是否已完成 |






<a name="warehouse.SpuStockReq"></a>

### SpuStockReq
spu库存请求结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spu_id | [int64](#int64) |  | spu_id |
| sku_ids | [int64](#int64) | repeated | sku_id |






<a name="warehouse.StockNumReply"></a>

### StockNumReply
库存数量响应结构


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 库存数量 |





 

 

 


<a name="warehouse.Warehouse"></a>

### Warehouse
内部服务
仓储服务接口定义

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSkuStock | [SkuStockReq](#warehouse.SkuStockReq) | [StockNumReply](#warehouse.StockNumReply) | 获取sku库存数量 |
| GetSpuStock | [SpuStockReq](#warehouse.SpuStockReq) | [SkuStockNum](#warehouse.SkuStockNum) | 获取spu下所有sku库存数量 |
| SKuStockLock | [SkuStockLockReq](#warehouse.SkuStockLockReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 锁定sku库存 |
| SkuStockUnlock | [SkuStockUnlockReq](#warehouse.SkuStockUnlockReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 解锁sku库存 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

