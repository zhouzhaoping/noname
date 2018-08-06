# -*- coding:utf-8 -*-
#
# Author: miaoyin
# Time: 2018/8/4 20:05

import requests
import json
import re
from bs4 import BeautifulSoup
import queue
import argparse


class StarSpider(object):
    def __init__(self, star_id):
        if star_id == 0:
            self.star_name = '易烊千玺'
            self.netease_id = 32
        if star_id == 1:
            self.star_name = '吴亦凡'
            self.netease_id = 21
        if star_id == 2:
            self.star_name = '蔡徐坤'
            self.netease_id = 82

    def _get_content(url):
        r = requests.get(url)
        r.encoding = 'utf-8'
        html = r.text
        soup = BeautifulSoup(html, 'html.parser')
        news_title = soup.select('article > div.head > h1')[0]
        news_content = soup.select('div.content > div.page')
        return news_title, news_content

    def _get_news_info(self, url):
        r = requests.get(url)
        html = r.text
        soup = BeautifulSoup(html, 'html.parser')
        news_infos = soup.select('div.info > span')
        news_time = news_infos[0].get_text()
        news_source = news_infos[1].get_text()

        return news_time, news_source

    def __call__(self, count):
        page = 0
        url = 'https://star.3g.163.com/star/article/list/{}-10.html?starId={}&callback='.format(page, self.netease_id)
        res = requests.get(url).text
        infos = json.loads(res)
        title_pattern = re.compile(self.star_name)
        news_list = []

        while page < 400:
            for info in infos['data']:
                if re.search(title_pattern, info['title']):
                    if len(info['pic_info']) == 0:
                        continue

                    cur_title = info['title']
                    cur_url = info['link']
                    cur_img = info['pic_info'][0]['url']
                    cur_create_time, cur_source = self._get_news_info(cur_url)
                    news_list.append({'title':cur_title, 'url':cur_url, 'img':cur_img, 'create_time':cur_create_time, 'source':cur_source})

                if len(news_list) == count:
                    return news_list

            page += 10

        return news_list


if __name__ == '__main__':
    arg_parser = argparse.ArgumentParser(description='manual to this script')
    arg_parser.add_argument('--starid', type=int, default=0)
    arg_parser.add_argument('--count', type=int, default=100)
    arg_parser.add_argument('--filename', type=str, default='news_list.json')
    args = arg_parser.parse_args()
    file_name = args.filename
    spider = StarSpider(args.starid)
    news_list = spider(args.count)

    print (json.dumps(news_list))
    #with open(file_name, 'w') as f:
    #    json.dump(news_list, f)







