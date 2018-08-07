# -*- coding:utf-8 -*-
# 
# Author: miaoyin
# Time: 2018/8/6 20:40

import requests
import json
import re
from bs4 import BeautifulSoup
import argparse
import time


class RankSpider(object):
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

    def _get_star_info(self, url):
        r = requests.get(url)
        html = r.text
        soup = BeautifulSoup(html, 'html.parser')
        news_infos = soup.select('div.info > span')
        news_time = news_infos[0].get_text()
        news_source = news_infos[1].get_text()

        return news_time, news_source

    def __call__(self, ):
        page = 0
        pattern = re.compile(self.star_name)

        while page < 300:
            url = 'https://star.3g.163.com/star/rank/list/{}-10.html?callback='.format(page)
            res = requests.get(url).text
            infos = json.loads(res)
            for info in infos['data']:
                if re.search(pattern, info['name']):
                    netease_daily_rank = info['currentRank']
                    star_rank = {'star_name':self.star_name, 'netease_daily_rank':netease_daily_rank}
                    return star_rank
            page += 10

        return {}


if __name__ == '__main__':
    arg_parser = argparse.ArgumentParser(description='manual to this script')
    arg_parser.add_argument('--starid', type=int, default=0)
    arg_parser.add_argument('--filename', type=str, default='star_rank.json')
    args = arg_parser.parse_args()
    file_name = args.filename
    spider = RankSpider(args.starid)
    star_rank = spider()

    with open(file_name, 'w') as f:
        json.dump(star_rank, f)
