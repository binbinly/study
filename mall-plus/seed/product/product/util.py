# 解析url地址
def parse_img_url(img_url):
    index = img_url.find('images')
    if index > -1:
        return img_url[index:]
    else:
        return img_url
