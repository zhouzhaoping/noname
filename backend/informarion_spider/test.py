#!/usr/bin/python
# Filename : helloworld.py
import time
import argparse
if __name__ == '__main__':
    arg_parser = argparse.ArgumentParser(description='manual to this script')
    arg_parser.add_argument('--starid', type=int, default=1)
    arg_parser.add_argument('--count', type=int, default=50)
    arg_parser.add_argument('--time', type=str, default='')
    arg_parser.add_argument('--filename', type=str, default='news_list.json')
    args = arg_parser.parse_args()
    print args
    print 'Hello World'
    #time.sleep(10)
    print 'fuck off'
