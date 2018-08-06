## 一、服务器
### 1.服务器配置
1. 服务器
	- ip:47.95.7.10
	- ssh端口:22/服务器端口:80
	- 用户名:root
	- 密码:请找周钊平同学私聊  
2. mysql
	- 用户名:root
	- 密码:请找周钊平同学私聊  
3. mysql远程账户
	- 端口:3306
	- 用户名:test
	- 密码:请找周钊平同学私聊      

|平台|CPU|内存|带宽|系统盘|操作系统|到期日期|价格|
|---|---|---|---|---|---|---|---|
|aliyun|Intel(R) Xeon(R) Platinum 8163 CPU @ 2.50GHz|2G内存|1M|40GB|CentOS 7.3|2019-02-04|57块钱半年|

防火墙当前打开的端口  
https://swas.console.aliyun.com
  
|应用类型|协议|端口范围|
|---|---|---|
|HTTP	|TCP	|80	|
|HTTPS	|TCP	|443|
|SSH	|TCP	|22	|
|MYSQL	|TCP	|3306|
|自定义	|TCP	|8080|

### 2.后台配置
方便日后可能的迁移：  
- go 1.10.3下载安装
	1. 下载`https://golang.org/doc/install?download=go1.10.3.linux-amd64.tar.gz`
	2. 解压到`/usr/local`
	3. /etc/profile添加：`export PATH=$PATH:/usr/local/go/bin`
	4. iris框架：`go get -u github.com/kataras/iris`
	5. 通过80端口访问

- mysql 5.7.23下载安装（https://www.jianshu.com/p/5dbabaf096f4）
	1. 启动`systemctl start mysqld`
	2. 启动状态`systemctl status mysqld`
	3. 开机启动`systemctl enable mysqld / systemctl daemon-reload`
	4. 登陆`mysql -uroot -p`
	5. 端口`3306`	
- go and mysql
	- sql驱动`go get github.com/go-sql-driver/mysql`
	- orm`go get github.com/go-xorm/xorm`
	- 随机数生成器`go get github.com/seehuhn/mt19937`
	- url路由接口`go get github.com/gorilla/mux`
- 文件夹
	- 项目位置`/pickme`
	- 图片保存位置`/srv/www/storage/image/`，详见`imagetool/config.json`
- python
	- `go get github.com/sbinet/go-python`
	- 需要先`sudo yum install python2-dev`

### 3.翻墙代理
ss代理:68.168.133.152:8385
密码:pickme
可翻墙供爬虫使用

## 二、接口格式
json