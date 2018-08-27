# -*- coding:utf-8 -*-
# 
# Author: miaoyin
# Time: 2018/8/4 20:05

import requests
import json
import re
from bs4 import BeautifulSoup
import argparse
import time
import datetime
import urllib


class StarSpider(object):
    def __init__(self, name, neteasyid):
        self.star_name = name
        self.netease_id = neteasyid

    def get_netease_starnews(self, count, update_time_stamp):

        def get_content(url):
            r = requests.get(url)
            r.encoding = 'utf-8'
            html = r.text
            soup = BeautifulSoup(html, 'html.parser')
            news_title = soup.select('article > div.head > h1')[0]
            news_content = soup.select('div.content > div.page')
            return news_title, news_content

        def get_news_info(url):
            r = requests.get(url)
            html = r.text
            soup = BeautifulSoup(html, 'html.parser')
            news_infos = soup.select('div.info > span')
            try:
                news_time = news_infos[0].get_text()
                news_source = news_infos[1].get_text()
            except:
                return None, None

            return news_time, news_source
        page = 0
        title_pattern = re.compile(self.star_name)
        netease_news_list = []

        while page < 400:
            url = 'https://star.3g.163.com/star/article/list/{}-10.html?starId={}&callback='.format(page, self.netease_id)
            res = requests.get(url).text
            infos = json.loads(res)
            if not infos:
                continue
            for info in infos['data']:
                cur_url = info['link']
                cur_create_time, cur_source = get_news_info(cur_url)
                if not cur_create_time:
                    continue
                time_array = time.strptime(cur_create_time, "%Y-%m-%d %H:%M:%S")
                time_stamp = int(time.mktime(time_array))
                if time_stamp < update_time_stamp:
                    return netease_news_list
                if re.search(title_pattern, info['title']):
                    if len(info['pic_info']) == 0:
                        continue

                    cur_title = info['title']
                    cur_img = info['pic_info'][0]['url']
                    netease_news_list.append({'title':cur_title, 'news_url':cur_url, 'img':cur_img, 'create_time':cur_create_time, 'source':cur_source})

                if len(netease_news_list) == count:
                    return netease_news_list

            page += 10

        return netease_news_list

    def get_netease_entnews(self, count, update_time_stamp):
        title_pattern = re.compile((self.star_name))
        netease_news_list = []

        for page in range(10):
            if page == 0:
                url = 'http://ent.163.com/special/000380VU/newsdata_index.js?callback='
            else:
                url = 'http://ent.163.com/special/000380VU/newsdata_index_{}.js?callback='.format(str(page+1).zfill(2))
            res = requests.get(url).text
            if not res:
                continue
            pattern = re.compile('^data_callback\(')
            res = re.sub(pattern, '', res)
            pattern = re.compile('\)$')
            res = re.sub(pattern, '', res)
            infos = json.loads(res)
            if not infos:
                continue
            for info in infos:
                cur_create_time = info['time'] #08/25/2018 14:02:54
                time_array = time.strptime(cur_create_time, '%m/%d/%Y %H:%M:%S')
                time_stamp = int(time.mktime(time_array))
                d = datetime.datetime.fromtimestamp(time_stamp)
                cur_create_time = d.strftime("%Y-%m-%d %H:%M:%S")
                if time_stamp < update_time_stamp:
                    return netease_news_list
                if re.search(title_pattern, info['title']):
                    if not info['imgurl']:
                        continue
                    cur_title = info['title']
                    cur_url = info['docurl']
                    cur_img = info['imgurl']
                    netease_news_list.append(
                        {'title': cur_title, 'news_url': cur_url, 'img': cur_img, 'create_time': cur_create_time,
                         'source': '网易娱乐'})
                if len(netease_news_list) == count:
                    return netease_news_list

        return netease_news_list

    def get_tencent_news(self, count, update_time_stamp):
        title_pattern = re.compile(self.star_name)
        tencent_news_list = []

        for page in range(15):
            url = 'https://pacaio.match.qq.com/irs/rcd?cid=58&token=c232b098ee7611faeffc46409e836360&ext=ent&page={}&expIds=&callback='.format(page)
            res = requests.get(url).text
            infos = json.loads(res)
            if not infos:
                continue
            for info in infos['data']:
                cur_create_time = info['publish_time']
                time_array = time.strptime(cur_create_time, "%Y-%m-%d %H:%M:%S")
                time_stamp = int(time.mktime(time_array))
                if time_stamp < update_time_stamp:
                    return tencent_news_list
                if re.search(title_pattern, info['title']):
                    if info['img'] == '':
                        continue
                    cur_title = info['title']
                    cur_url = info['vurl']
                    cur_img = info['img']
                    tencent_news_list.append({'title':cur_title, 'news_url':cur_url, 'img':cur_img, 'create_time':cur_create_time, 'source':'腾讯娱乐'})

                if len(tencent_news_list) == count:
                    return tencent_news_list

        return tencent_news_list

    def get_sina_news(self, count, update_time_stamp):
        sina_news_list = []
        title_pattern = re.compile(self.star_name)
        time_stamp = int(time.time())
        for i in range(20):
            url = 'http://feed.mix.sina.com.cn/api/roll/get?pageid=107&lid=1244&num={}&versionNumber=1.2.8&ctime={}&encode=utf-8&callback='.format(30, time_stamp)
            # req = urllib.request.Request(url)
            # req.add_header('User-Agent', 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36')
            # res = urllib.request.urlopen(req).read().decode('utf-8')
            res = requests.get(url).text
            if not res:
                continue
            infos = json.loads(res)
            if not infos:
                continue
            for info in infos['result']['data']:
                time_stamp = int(info['ctime'])
                if time_stamp < update_time_stamp:
                    return sina_news_list
                d = datetime.datetime.fromtimestamp(time_stamp)
                cur_create_time = d.strftime("%Y-%m-%d %H:%M:%S")
                if re.search(title_pattern, info['title']):
                    cur_title = info['title']
                    if int(info['ctime']) < update_time_stamp:
                        return sina_news_list
                    if 'u' not in info['img']:
                        continue
                    cur_img = info['img']['u']
                    cur_url = info['url']
                    sina_news_list.append({'title':cur_title, 'news_url':cur_url, 'img':cur_img, 'create_time':cur_create_time, 'source':'新浪娱乐'})

                if len(sina_news_list) == count:
                    return sina_news_list

        return sina_news_list

    def get_cnnewsent_news(self, count, update_time_stamp):
        cn_news_list = []
        title_pattern = re.compile(self.star_name)

        for i in range(5):
            url = 'http://channel.chinanews.com/cns/s/channel:yl.shtml?pager={}&pagenum={}'.format(i, 100)
            res = requests.get(url).text.replace('var specialcnsdata = ', '')
            text_pattern = re.compile(';.$')
            res = re.sub(text_pattern, '', res)
            infos = json.loads(res)
            if not infos:
                continue
            for info in infos['docs']:
                cur_create_time = info['pubtime'] + ':00'
                time_array = time.strptime(cur_create_time, "%Y-%m-%d %H:%M:%S")
                time_stamp = int(time.mktime(time_array))
                if time_stamp < update_time_stamp:
                    return cn_news_list
                if re.search(title_pattern, info['title']):
                    cur_title = info['title']
                    cur_img = info['img_cns']
                    if not cur_img:
                        continue
                    cur_url = info['url']
                    cn_news_list.append({'title':cur_title, 'news_url':cur_url, 'img':cur_img, 'create_time':cur_create_time, 'source':'中国新闻网'})

                    if len(cn_news_list) == count:
                        return cn_news_list
        return cn_news_list

    def __call__(self, count, update_time):

        assert self.star_name, 'Star name is necessary'

        if len(update_time):
            update_time_array = time.strptime(update_time, "%Y-%m-%d %H:%M:%S")
            update_time_stamp = int(time.mktime(update_time_array))
        else:
            update_time_stamp = 0

        news_list = []

        if self.netease_id:
            news_list.extend(self.get_netease_starnews(count, update_time_stamp))
        else:
            news_list.extend(self.get_netease_entnews(count, update_time_stamp))
        news_list.extend(self.get_tencent_news(count, update_time_stamp))
        news_list.extend(self.get_sina_news(count, update_time_stamp))
        news_list.extend(self.get_cnnewsent_news(count, update_time_stamp))

        def take_ctime(elem):
            return int(time.mktime(time.strptime(elem['create_time'], "%Y-%m-%d %H:%M:%S")))
        news_list.sort(key=take_ctime, reverse=True)
        news_list = news_list[:count]

        return news_list


if __name__ == '__main__':
    arg_parser = argparse.ArgumentParser(description='manual to this script')
    arg_parser.add_argument('--neteasyid', type=str, default='')
    arg_parser.add_argument('--name', type=str, default='')
    arg_parser.add_argument('--count', type=int, default=50)
    arg_parser.add_argument('--time', type=str, default='')
    arg_parser.add_argument('--filename', type=str, default='news_list.json')
    args = arg_parser.parse_args()
    file_name = args.filename
    spider = StarSpider(args.name, args.neteasyid)
    news_list = spider(args.count, args.time)

    with open(file_name, 'w') as f:
        json.dump(news_list, f)

    print(json.dumps(news_list))







