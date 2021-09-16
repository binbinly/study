import scrapy


class WydlpSpider(scrapy.Spider):
    name = 'wydlp'
    allowed_domains = ['wydlp.com']
    start_urls = ['https://www.wydlp.com/']

    def parse(self, response):
        pass
