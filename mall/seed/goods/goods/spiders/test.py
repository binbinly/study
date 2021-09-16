import scrapy
import time
import re
import itertools
import json
from goods.items import *


class TestSpider(scrapy.Spider):
    name = 'test'
    allowed_domains = ['home72.com']
    start_urls = ['http://www.home72.com/goods.php?id=925',
                  'http://www.home72.com/goods.php?id=1280',
                  'http://www.home72.com/goods.php?id=1537',
                  'http://www.home72.com/goods.php?id=1283',
                  'http://www.home72.com/goods.php?id=268', ]
    host = 'http://www.home72.com'
    attr_name_id = 1  # 销售属性主键
    attr_val_id = 1  # 销售属性值主键
    attrs = {}
    attr_val = {}

    def parse(self, response):
        r = re.search(r'php\?id=(.*)', response.url, re.M | re.I)
        goods_id = r.group(1)

        # 商品 sku
        attrs = []
        attr_names = []
        value_names = {}
        for quote in response.xpath(
                '//div[@id="goodsInfo"]//ul[@class="fd30_ulinfo"]//li[@class="padd loop fd30_ulinfo"]'):
            attr_name = quote.xpath('strong/text()').get().replace('：', '')
            attr_name_id = self.attr_name_id
            if attr_name not in self.attrs:
                sku_attr = SkuAttrItem()
                sku_attr['name'] = attr_name
                sku_attr['id'] = attr_name_id
                yield sku_attr
            else:
                attr_name_id = self.attrs[attr_name]
            attr_names.append(attr_name_id)

            attr = []
            attr_values = []
            for n in quote.xpath('div[1]/input/@dir').extract():
                attr_value_id = self.attr_val_id
                if n not in self.attr_val:
                    sku_attr_val = SkuAttrValItem()
                    sku_attr_val['id'] = attr_value_id
                    sku_attr_val['value'] = n
                    sku_attr_val['attr_id'] = attr_name_id
                    yield sku_attr_val
                else:
                    attr_value_id = self.attr_val[n]

                attr_values.append({'id': attr_value_id, 'name': n})
                value_names[attr_value_id] = n
                attr.append(attr_value_id)
                self.attr_val_id += 1

            goods_sku_attr = GoodsSkuAttrItem()
            goods_sku_attr['goods_id'] = goods_id
            goods_sku_attr['attr_id'] = attr_name_id
            goods_sku_attr['attr_name'] = attr_name
            goods_sku_attr['values'] = json.dumps(attr_values)
            yield goods_sku_attr

            attrs.append(attr)
            self.attr_name_id += 1

        # 商品
        goods = GoodsItem()
        goods['id'] = goods_id
        goods['title'] = response.xpath('//div[@id="goodsInfo"]//div[@class="goodsnames"]/text()').get()
        goods['intro'] = response.xpath('//div[@id="goodsInfo"]//div[@class="briefs"]/text()').get()
        goods['price'] = response.xpath('//div[@id="goodsInfo"]//font[@id="ECS_GOODS_AMOUNT"]/text()').re_first(
            r'￥(.*)')
        goods['original_price'] = response.xpath(
            '//div[@id="goodsInfo"]//div[@class="all_price"]//div[3]/font/text()').re_first(r'￥(.*)')
        goods['cover'] = self.start_urls[0] + response.xpath(
            '//div[@id="goodsInfo"]//div[@class="gmax_pic_box"]//img/@src').get()
        goods['stock'] = 100
        goods['cat_id'] = response.xpath('//div[@id="ur_here_g"]/div/a[last()]/@href').re_first(
            r'php\?id\=(.*)')
        goods['unit'] = '个'
        goods['created_at'] = int(time.time())
        goods['updated_at'] = int(time.time())
        goods['sku_many'] = 1
        goods['price'] = int(float(goods['price']) * 100)
        goods['original_price'] = int(float(goods['original_price']) * 100)
        if len(attrs) == 0:  # 单规格
            goods['sku_many'] = 0
        # yield goods

        # 计算笛卡尔积
        sku_attrs = ','.join([str(v) for v in attr_names])
        for item in itertools.product(*attrs):
            sku = GoodsSkuItem()
            sku['goods_id'] = goods_id
            sku['stock'] = 100
            sku['price'] = goods['price']
            sku['original_price'] = goods['original_price']
            sku['values'] = ','.join([str(v) for v in item])
            sku['attrs'] = sku_attrs
            names = []
            for value_id in item:
                if value_id in value_names:
                    names.append(value_names[value_id])
            sku['value_names'] = ','.join(names)
            yield sku
