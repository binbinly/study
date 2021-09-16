# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html
from scrapy import Field, Item


# 商品结构
class GoodsItem(Item):
    id = Field()
    title = Field()
    cover = Field()
    cat_id = Field()
    price = Field()
    original_price = Field()
    intro = Field()
    stock = Field()
    sku_many = Field()
    unit = Field()
    created_at = Field()
    updated_at = Field()


# 商品分类
class GoodsCategoryItem(Item):
    id = Field()
    pid = Field()
    name = Field()
    created_at = Field()
    updated_at = Field()


# 商品属性
class GoodsAttrItem(Item):
    goods_id = Field()
    name = Field()
    value = Field()


# 商品图片
class GoodsImageItem(Item):
    goods_id = Field()
    type = Field()
    url = Field()
    created_at = Field()


# 商品SKU
class GoodsSkuItem(Item):
    goods_id = Field()
    stock = Field()
    price = Field()
    original_price = Field()
    attrs = Field()
    values = Field()
    value_names = Field()


# 商品销售属性
class GoodsSkuAttrItem(Item):
    goods_id = Field()
    attr_id = Field()
    attr_name = Field()
    values = Field()


# 销售属性
class SkuAttrItem(Item):
    id = Field()
    name = Field()
    desc = Field()


# 销售属性值
class SkuAttrValItem(Item):
    id = Field()
    attr_id = Field()
    value = Field()
    desc = Field()
