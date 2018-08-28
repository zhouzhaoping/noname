#!/usr/bin/env python
# -*- coding: UTF-8 -*-

import os
import re
import requests
import traceback
import json
import argparse
import time
import sys
#import click
import codecs
reload(sys)
sys.setdefaultencoding('gbk')
from datetime import datetime
from datetime import timedelta
from lxml import etree


class Weibo:
    cookie = {"Cookie": "_T_WM=b6f6b815e19c0f1c0949e1958cce0231; WEIBOCN_FROM=1110006030; SUB=_2A252bGFeDeRhGeRH6FcY9yjKzjmIHXVVrw8WrDV6PUJbkdAKLW_AkW1NTau4YJL-9AV6Rt8hSLOWD1DHYicFVvzm; SUHB=0ebdfs-sFIvmip; SCF=AhcMP28AzQgkknxBWeMkxq_bcOkq2JMcTmL6dLQFkvN0CUVUIoAAM4riSVN4oBQLGPMSRNxxwMq1jZcxIh8MArA.; SSOLoginState=1533546766; MLOGIN=1; M_WEIBOCN_PARAMS=lfid%3D1076033623353053%26luicode%3D20000174%26uicode%3D20000174"}  # 将your cookie替换成自己的cookie

    # Weibo类初始化
    def __init__(self, user_id, filter=0):
        self.user_id = user_id  # 用户id，即需要我们输入的数字，如昵称为“Dear-迪丽热巴”的id为1669879400
        self.filter = filter  # 取值范围为0、1，程序默认值为0，代表要爬取用户的全部微博，1代表只爬取用户的原创微博
        self.username = ''  # 用户名，如“Dear-迪丽热巴”
        self.weibo_num = 0  # 用户全部微博数
        self.weibo_num2 = 0  # 爬取到的微博数
        self.following = 0  # 用户关注数
        self.followers = 0  # 用户粉丝数
        self.weibo_content = []  # 微博内容
        self.weibo_place = []  # 微博位置
        self.publish_time = []  # 微博发布时间
        self.up_num = []  # 微博对应的点赞数
        self.retweet_num = []  # 微博对应的转发数
        self.comment_num = []  # 微博对应的评论数
        self.publish_tool = []  # 微博发布工具
        self.weibo_img = [] # 微博图片
        self.source = [] #来源
    # 获取用户昵称
    def get_username(self):
        try:
            url = "https://weibo.cn/%s/info" % (self.user_id)
            html = requests.get(url, cookies=self.cookie).content
            selector = etree.HTML(html)
            username = selector.xpath("//title/text()")[0]
            self.username = username[:-3]
        except Exception, e:
            print "Error: ", e
            traceback.print_exc()

    # 获取用户微博数、关注数、粉丝数
    def get_user_info(self):
        try:
            url = "https://weibo.cn/u/%s?filter=%d&page=1" % (
                self.user_id, self.filter)
            html = requests.get(url, cookies=self.cookie).content
            selector = etree.HTML(html)
            pattern = r"\d+\.?\d*"

            # 微博数
            str_wb = selector.xpath(
                "//div[@class='tip2']/span[@class='tc']/text()")[0]
            guid = re.findall(pattern, str_wb, re.S | re.M)
            for value in guid:
                num_wb = int(value)
                break
            self.weibo_num = num_wb

            # 关注数
            str_gz = selector.xpath("//div[@class='tip2']/a/text()")[0]
            guid = re.findall(pattern, str_gz, re.M)
            self.following = int(guid[0])

            # 粉丝数
            str_fs = selector.xpath("//div[@class='tip2']/a/text()")[1]
            guid = re.findall(pattern, str_fs, re.M)
            self.followers = int(guid[0])

        except Exception, e:
            print "Error: ", e
            traceback.print_exc()


    # 获取用户微博内容及对应的发布时间、点赞数、转发数、评论数
    def get_weibo_info(self):
        try:
            url = "https://weibo.cn/u/%s?filter=%d&page=1" % (
                self.user_id, self.filter)
            html = requests.get(url, cookies=self.cookie).content
            selector = etree.HTML(html)
            if selector.xpath("//input[@name='mp']") == []:
                page_num = 1
            else:
                page_num = (int)(selector.xpath(
                    "//input[@name='mp']")[0].attrib["value"])
            pattern = r"\d+\.?\d*"
            for page in range(1, page_num + 1):
                url2 = "https://weibo.cn/u/%s?filter=%d&page=%d" % (
                    self.user_id, self.filter, page)
                html2 = requests.get(url2, cookies=self.cookie).content
                selector2 = etree.HTML(html2)
                info = selector2.xpath("//div[@class='c']")
                is_empty = info[0].xpath("div/span[@class='ctt']")
                if is_empty:
                    for i in range(0, len(info) - 2):
                        # 微博内容
                        str_t = info[i].xpath("div/span[@class='ctt']")
                        weibo_content = str_t[0].xpath("string(.)").encode(
                            sys.stdout.encoding, "ignore").decode(
                            sys.stdout.encoding)
                        weibo_content = weibo_content[:-1]
                        weibo_id = info[i].xpath("@id")[0][2:]
                        a_link = info[i].xpath(
                            "div/span[@class='ctt']/a/@href")
                        if a_link:
                            if (a_link[-1] == "/comment/" + weibo_id or
                                    "/comment/" + weibo_id + "?" in a_link[-1]):
                                weibo_link = "https://weibo.cn" + a_link[-1]
                                wb_content = self.get_long_weibo(weibo_link)
                                if wb_content:
                                    weibo_content = wb_content
                        self.weibo_content.append(weibo_content)

                        #图片
                        weibo_img = u"无"
                        if len(info[i].xpath("div")) > 1:
                            div_first = info[i].xpath("div")[1]
                            img_list = div_first.xpath("a")
                            for img in img_list:
                                if img.xpath("img/@src"):
                                    weibo_img = img.xpath("img/@src")[-1].encode(
                                        sys.stdout.encoding, "ignore").decode(sys.stdout.encoding)
                                    break
                        self.weibo_img.append(weibo_img)

                        # 微博位置
                        div_first = info[i].xpath("div")[0]
                        a_list = div_first.xpath("a")
                        weibo_place = u"无"
                        for a in a_list:
                            if ("http://place.weibo.com/imgmap/center" in a.xpath("@href")[0] and
                                    a.xpath("text()")[0] == u"显示地图"):
                                weibo_place = div_first.xpath(
                                    "span[@class='ctt']/a")[-1]
                                if u"的秒拍视频" in div_first.xpath("span[@class='ctt']/a/text()")[-1]:
                                    weibo_place = div_first.xpath(
                                        "span[@class='ctt']/a")[-2]
                                weibo_place = weibo_place.xpath("string(.)").encode(
                                    sys.stdout.encoding, "ignore").decode(sys.stdout.encoding)
                                break
                        self.weibo_place.append(weibo_place)

                        # 微博发布时间
                        str_time = info[i].xpath("div/span[@class='ct']")
                        str_time = str_time[0].xpath("string(.)").encode(
                            sys.stdout.encoding, "ignore").decode(
                            sys.stdout.encoding)
                        publish_time = str_time.split(u'来自')[0]
                        if u"刚刚" in publish_time:
                            publish_time = datetime.now().strftime(
                                '%Y-%m-%d %H:%M')
                        elif u"分钟" in publish_time:
                            minute = publish_time[:publish_time.find(u"分钟")]
                            minute = timedelta(minutes=int(minute))
                            publish_time = (
                                datetime.now() - minute).strftime(
                                "%Y-%m-%d %H:%M")
                        elif u"今天" in publish_time:
                            today = datetime.now().strftime("%Y-%m-%d")
                            time = publish_time[3:]
                            publish_time = today + " " + time
                        elif u"月" in publish_time:
                            year = datetime.now().strftime("%Y")
                            month = publish_time[0:2]
                            day = publish_time[3:5]
                            time = publish_time[7:12]
                            publish_time = (
                                year + "-" + month + "-" + day + " " + time)
                        else:
                            publish_time = publish_time[:16]
                        self.publish_time.append(publish_time)

                        # 微博发布工具
                        if len(str_time.split(u'来自')) > 1:
                            publish_tool = str_time.split(u'来自')[1]
                        else:
                            publish_tool = u"无"
                        self.publish_tool.append(publish_tool)

                        str_footer = info[i].xpath("div")[-1]
                        str_footer = str_footer.xpath("string(.)").encode(
                            sys.stdout.encoding, "ignore").decode(sys.stdout.encoding)
                        str_footer = str_footer[str_footer.rfind(u'赞'):]
                        guid = re.findall(pattern, str_footer, re.M)

                        # 点赞数
                        up_num = int(guid[0])
                        self.up_num.append(up_num)

                        # 转发数
                        retweet_num = int(guid[1])
                        self.retweet_num.append(retweet_num)

                        # 评论数
                        comment_num = int(guid[2])
                        self.comment_num.append(comment_num)

                        self.weibo_num2 += 1

        except Exception, e:
            print "Error: ", e
            traceback.print_exc()

    # 将爬取的信息写入文件
    def write_txt(self, update_time):
        try:
            if self.filter:
                result_header = u"\n\n原创微博内容: \n"
            else:
                result_header = u"\n\n微博内容: \n"
            result = (u"用户信息\n用户昵称：" + self.username +
                      u"\n用户id: " + str(self.user_id) +
                      u"\n微博数: " + str(self.weibo_num) +
                      u"\n关注数: " + str(self.following) +
                      u"\n粉丝数: " + str(self.followers) +
                      result_header
                      )
            if len(update_time):
                update_time_array = time.strptime(update_time, "%Y-%m-%d %H:%M")
                update_time_stamp = int(time.mktime(update_time_array))
            else:
                update_time_stamp = 0

            result = []
            for i in range(1, self.weibo_num2 + 1):
                result.append({"account_id": self.user_id,
                        "account_name": self.username,
                        "content": self.weibo_content[i - 1],
                        "create_time": self.publish_time[i - 1],
                        "imgs": self.weibo_img[i - 1],
                        "source": "weibo"
                        })
            print json.dumps(result)
            file_dir = os.path.split(os.path.realpath(__file__))[
                0] + os.sep + "weibo"
            if not os.path.isdir(file_dir):
                os.mkdir(file_dir)
            file_path = file_dir + os.sep + "%s" % self.user_id + "_" + "%s" % self.username + ".json"
            with open(file_path, 'w') as f:
                json.dump(result, f)
        except Exception, e:
            print "Error: ", e
            traceback.print_exc()

    # 运行爬虫
    def start(self, update_time):
        try:
            self.get_username()
            self.get_weibo_info()
            self.write_txt(update_time)
        except Exception, e:
            print "Error: ", e

class Instagram:
    def __init__(self, user_id):
        self.user_id = user_id
        print user_id
        self.BASE_URL = "https://www.instagram.com/" + self.user_id + "/"
        print self.BASE_URL
        self.headers = {
            "Origin": "https://www.instagram.com/",
            "Referer": self.BASE_URL,
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
                          "Chrome/58.0.3029.110 Safari/537.36",
            "Host": "www.instagram.com",
            "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
            "accept-encoding": "gzip, deflate, sdch, br",
            "accept-language": "zh-CN,zh;q=0.8",
            "X-Instragram-AJAX": "1",
            "X-Requested-With": "XMLHttpRequest",
            "Upgrade-Insecure-Requests": "1",
        }

        self.top_url = ''
        self.account_id = ''
        self.account_name = self.user_id
        self.contents = []
        self.ins_num = 0  # 爬取到的ins数
        self.new_imgs_url = []
        self.create_time = []
    def crawl(self, update_time):
        #click.echo('start')
        try:
            res = requests.get(self.BASE_URL, headers=self.headers)
            #print "here"
            html = etree.HTML(res.content.decode('utf-8'))
            #print html
            all_a_tags = html.xpath('//script[@type="text/javascript"]/text()')

            for a_tag in all_a_tags:
                #print "a_tag:" + a_tag
                if a_tag.strip().startswith('window._sharedData'):
                    data = a_tag.split('= {')[1][:-1]  # 获取json数据块
                    js_data = json.loads('{' + data, encoding='utf-8')
                    user = js_data["entry_data"]["ProfilePage"][0]["graphql"]["user"]
                    self.account_id = user["id"]
                    edges = user["edge_owner_to_timeline_media"]["edges"]
                    for edge in edges:
                        if self.top_url and self.top_url == edge["node"]["display_url"]:
                            in_top_url_flag = True
                            break
                        #click.echo(edge["node"]["display_url"])
                        self.new_imgs_url.append(edge["node"]["display_url"])
                        content_candidate = edge["node"]["edge_media_to_caption"]["edges"]
                        if len(content_candidate) > 0:
                            self.contents.append(content_candidate[0]["node"]["text"])
                            print "content:" + content_candidate[0]["node"]["text"]
                        else:
                            self.contents.append("")
                        self.create_time.append(time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(edge["node"]["taken_at_timestamp"])))
                    #click.echo('ok')

            self.write_txt(update_time)
        except Exception as e:
            raise e

    def write_txt(self, update_time):
        try:
            if len(update_time):
                update_time_array = time.strptime(update_time, "%Y-%m-%d %H:%M")
                update_time_stamp = int(time.mktime(update_time_array))
            else:
                update_time_stamp = 0

            result = []
            for i in range(1, len(self.new_imgs_url) + 1):
                time_array = time.strptime(self.create_time[i - 1], "%Y-%m-%d %H:%M:%S")
                time_stamp = int(time.mktime(time_array))
                if time_stamp < update_time_stamp:
                    break
                result.append({"account_id": self.account_id,
                        "account_name": self.account_name,
                        "content": self.contents[i - 1],
                        "create_time": self.create_time[i - 1],
                        "imgs": self.new_imgs_url[i - 1],
                        "source": u"Instagram"
                        })
            print json.dumps(result)
            file_dir = os.path.split(os.path.realpath(__file__))[
                0] + os.sep + "ins"
            if not os.path.isdir(file_dir):
                os.mkdir(file_dir)
            file_path = file_dir + os.sep + "%s" % self.account_id + "_" + "%s" % self.account_name + ".json"
            # f = open(file_path, "wb")
            # f.write(text.encode(sys.stdout.encoding))
            # f.close()
            with open(file_path, 'w') as f:
                json.dump(result, f)
        except Exception, e:
            print "Error: ", e
            traceback.print_exc()

#def main():
 #   try:
        # 使用实例,输入一个用户id，所有信息都会存储在wb实例中
        # user_id = 3623353053  # 可以改成任意合法的用户id（爬虫的微博id除外）
        # filter = 0  # 值为0表示爬取全部微博（原创微博+转发微博），值为1表示只爬取原创微博
        # wb = Weibo(user_id, filter)  # 调用Weibo类，创建微博实例wb
        # wb.start()  # 爬取微博信息
        # print u"用户名: " + wb.username
        # print u"全部微博数: " + str(wb.weibo_num)
        # print u"关注数: " + str(wb.following)
        # print u"粉丝数: " + str(wb.followers)
        # if wb.weibo_content:
        #     print u"最新/置顶 微博为: " + wb.weibo_content[0]
        #     print u"最新/置顶 微博位置: " + wb.weibo_place[0]
        #     print u"最新/置顶 微博发布时间: " + wb.publish_time[0]
        #     print u"最新/置顶 微博获得赞数: " + str(wb.up_num[0])
        #     print u"最新/置顶 微博获得转发数: " + str(wb.retweet_num[0])
        #     print u"最新/置顶 微博获得评论数: " + str(wb.comment_num[0])
        #     print u"最新/置顶 微博发布工具: " + wb.publish_tool[0]
    # except Exception, e:
    #     print "Error: ", e
    #     traceback.print_exc()


if __name__ == "__main__":
    arg_parser = argparse.ArgumentParser(description='manual to this script')
    arg_parser.add_argument('--user_id', type=str, default='3623353053') #默认是易烊千玺的微博id
    arg_parser.add_argument('--time', type=str, default='')
    arg_parser.add_argument('--source', type=str, default='')
    args = arg_parser.parse_args()
    filter = 0  # 值为0表示爬取全部微博（原创微博+转发微博），值为1表示只爬取原创微博
    if args.source == "weibo":
        wb = Weibo(args.user_id, filter)  # 调用Weibo类，创建微博实例wb
        wb.start(args.time)  # 爬取微博信息
    else:
        if args.source == "instagram":
            ins = Instagram(args.user_id)
            ins.crawl(args.time)
        else:
            print "missing parameters"


