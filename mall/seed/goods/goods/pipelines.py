# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import os
import hashlib
import datetime
import pymysql
from goods.items import *
import logging
import scrapy
from scrapy.pipelines.images import ImagesPipeline
from scrapy.utils.python import to_bytes


class MinioPipeline(ImagesPipeline):

    def __init__(self, store_uri, download_func=None, settings=None):
        self.sub_dir = datetime.date.today().strftime('%Y%m%d')
        super().__init__(store_uri, download_func, settings)

    def get_media_requests(self, item, info):
        if item.get('url'):
            return scrapy.Request(item["url"])

    def item_completed(self, results, item, info):
        image_paths = [x["path"] for ok, x in results if ok]
        if image_paths:
            item["url"] = image_paths[0]
        return item

    def file_path(self, request, response=None, info=None, *, item=None):
        image_guid = hashlib.sha1(to_bytes(request.url)).hexdigest()
        return f'images/{self.sub_dir}/{image_guid}.jpg'


class GoodsPipeline:

    def __init__(self, db_host, db_port, db_user, db_pwd, db_name):
        self.sub_dir = datetime.date.today().strftime('%Y%m%d')
        self.db_host = db_host
        self.db_port = db_port
        self.db_user = db_user
        self.db_pwd = db_pwd
        self.db_name = db_name

    @classmethod
    def from_crawler(cls, crawler):
        return cls(
            db_host=os.environ.get("MYSQL_HOST", "192.168.8.76"),
            db_pwd=os.environ.get("MYSQL_PWD", "123456"),
            db_name=os.environ.get("MYSQL_DB_NAME", "mall"),
            db_port=os.environ.get("MYSQL_PORT", "3306"),
            db_user=os.environ.get("MYSQL_USER", "root")
        )

    def open_spider(self, spider):
        self.db = pymysql.connect(
            host=self.db_host,
            user=self.db_user,
            password=self.db_pwd,
            db=self.db_name,
            charset="utf8mb4",
            cursorclass=pymysql.cursors.DictCursor,
        )
        self.cursor = self.db.cursor()

    def close_spider(self, spider):
        self.db.close()

    def process_item(self, item, spider):
        try:
            if isinstance(item, GoodsItem):
                image_guid = hashlib.sha1(to_bytes(item['cover'])).hexdigest()
                item['cover'] = f'images/{self.sub_dir}/{image_guid}.jpg'
                self.send_db(item, 'goods', item['id'])
            elif isinstance(item, GoodsCategoryItem):
                self.send_db(item, 'goods_category', item['id'])
            elif isinstance(item, GoodsAttrItem):
                sql = 'select * from goods_attr where goods_id=%s and `name`="%s"' % (item['goods_id'], item['name'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'goods_attr')
            elif isinstance(item, GoodsImageItem):
                sql = 'select * from goods_image where goods_id=%s and `url`="%s"' % (item['goods_id'], item['url'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'goods_image')
            elif isinstance(item, GoodsSkuItem):
                sql = 'select * from goods_sku where goods_id=%s and `values`="%s"' % (item['goods_id'], item['values'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'goods_sku')
            elif isinstance(item, GoodsSkuAttrItem):
                sql = 'select * from goods_sku_attr where goods_id=%s and `attr_id`=%s' % (
                    item['goods_id'], item['attr_id'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'goods_sku_attr')
            elif isinstance(item, SkuAttrItem):
                sql = 'select * from sku_attr where  `name`="%s"' % item['name']
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'sku_attr')
            elif isinstance(item, SkuAttrValItem):
                sql = 'select * from sku_attr_value where  attr_id=%s and `value`="%s"' % (
                    item['attr_id'], item['value'])
                self.cursor.execute(sql)
                if not self.cursor.fetchone():
                    self.save_db(item, 'sku_attr_value')
        except Exception as e:
            logging.error(item)
            logging.error(e)
        return item

    def send_db(self, item, table, pri_id):
        sql = 'select * from %s where id=%s' % (table, pri_id)
        self.cursor.execute(sql)
        if not self.cursor.fetchone():
            return self.save_db(item, table)

    def save_db(self, item, table):
        keys = item.keys()
        values = tuple(item.values())
        fields = ",".join(['`' + v + '`' for v in keys])
        temp = ",".join(["%s"] * len(keys))
        sql = "INSERT INTO {} ({}) VALUES ({})".format(table, fields, temp)
        self.cursor.execute(sql, values)
        return self.db.commit()
