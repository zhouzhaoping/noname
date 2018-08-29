/*
Navicat MySQL Data Transfer

Source Server         : pickme
Source Server Version : 50723
Source Host           : 47.95.7.10:3306
Source Database       : startsuck

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-08-12 11:36:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `agenda`
-- ----------------------------
DROP TABLE IF EXISTS `agenda`;
CREATE TABLE `agenda` (
  `agenda_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一自增',
  `star_id` int(11) NOT NULL COMMENT '明星id，引用star_id',
  `detail_time` datetime NOT NULL COMMENT '行程的详细时间',
  `location` varchar(255) NOT NULL COMMENT '行程的地点',
  `content` varchar(255) NOT NULL COMMENT '行程内容',
  PRIMARY KEY (`agenda_id`),
  KEY `agenda_star_id` (`star_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of agenda
-- ----------------------------
INSERT INTO `agenda` VALUES ('1', '0', '2018-08-24 00:00:00', '北京', 'TFBOYS五周年演唱会');
INSERT INTO `agenda` VALUES ('2', '1', '2018-08-26 00:00:00', '加拿大多伦多', '2018 iHeartRadio MMVAS');
INSERT INTO `agenda` VALUES ('3', '2', '2018-09-21 00:00:00', '无锡：华莱坞影视城', '《幻乐之城》总决赛录制（暂定）');

-- ----------------------------
-- Table structure for `auth_accounts`
-- ----------------------------
DROP TABLE IF EXISTS `auth_accounts`;
CREATE TABLE `auth_accounts` (
  `user_id` int(11) NOT NULL COMMENT '用户id,引用用户表的user_id',
  `user_name` varchar(255) NOT NULL COMMENT '授权系统的账户名，比如微博的用户名',
  `password` varchar(255) NOT NULL COMMENT '授权系统的账户密码，比如微博账户的密码',
  `account_type` int(11) NOT NULL COMMENT '账户类型，0：百度，1：微博，2：ins'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户授权的账户表';

-- ----------------------------
-- Records of auth_accounts
-- ----------------------------

-- ----------------------------
-- Table structure for `info_source`
-- ----------------------------
DROP TABLE IF EXISTS `info_source`;
CREATE TABLE `info_source` (
  `info_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '信息id，唯一主键，自增',
  `star_id` int(11) NOT NULL COMMENT '明星id，引用star_id',
  `source` varchar(255) DEFAULT NULL COMMENT '爬取链接，网易/ins/百度等',
  `account_name` varchar(255) DEFAULT NULL COMMENT '各来源的账号唯一标识符',
  `usage_type` int(11) DEFAULT NULL COMMENT '信息的用途类型，1:state,2:news',
  PRIMARY KEY (`info_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of info_source
-- ----------------------------

-- ----------------------------
-- Table structure for `login_log`
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log` (
  `suv` varchar(255) NOT NULL COMMENT '用户的suv',
  `login_time` datetime NOT NULL COMMENT '访问（登录）时间',
  `ip` varchar(255) DEFAULT NULL COMMENT '用户登录ip地址'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户浏览（登录）记录';

-- ----------------------------
-- Records of login_log
-- ----------------------------

-- ----------------------------
-- Table structure for `news`
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `news_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '资讯id，唯一自增',
  `star_id` int(11) NOT NULL COMMENT '明星id，引用',
  `img` varchar(255) NOT NULL COMMENT '资讯图片',
  `title` varchar(255) NOT NULL COMMENT '资讯标题',
  `news_url` varchar(255) DEFAULT NULL COMMENT '资讯链接',
  `source` varchar(255) NOT NULL DEFAULT '' COMMENT '资讯来源',
  `create_time` datetime DEFAULT NULL COMMENT '资讯创建时间',
  PRIMARY KEY (`news_id`),
  KEY `news_star_id` (`star_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of news
-- ----------------------------

-- ----------------------------
-- Table structure for `post`
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `post_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一自增',
  `user_id` int(10) NOT NULL COMMENT '用户id，引用user_id',
  `create_time` datetime NOT NULL COMMENT '帖子/评论创建时间',
  `title` varchar(255) NOT NULL COMMENT '帖子和评论的标题',
  `content` varchar(255) NOT NULL COMMENT '帖子和评论的内容',
  `like_num` int(10) unsigned zerofill NOT NULL DEFAULT '0000000000' COMMENT '帖子/评论点赞数',
  `star_id` int(11) DEFAULT NULL COMMENT '明星id,引用star_id，只有level 0的post该字段值不空',
  `parent_comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '父评论id，当该值为0时表示是一条帖子而不是评论',
  `comment_num` int(11) NOT NULL DEFAULT '0' COMMENT '帖子/评论的评论数',
  `level` int(11) NOT NULL COMMENT '级别：0：帖子，1：一级评论，2：二级评论',
  `imgs` varchar(255) DEFAULT NULL COMMENT '帖子的图片，评论没有图片，所有只有当level=0时该字段不空',
  PRIMARY KEY (`post_id`),
  KEY `post_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帖子和评论共用一张表';

-- ----------------------------
-- Records of post
-- ----------------------------

-- ----------------------------
-- Table structure for `post_user_relation`
-- ----------------------------
DROP TABLE IF EXISTS `post_user_relation`;
CREATE TABLE `post_user_relation` (
  `post_id` int(11) NOT NULL COMMENT '帖子/评论id',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `is_like` int(11) NOT NULL COMMENT '用户对该帖子/评论是否点赞，1：点赞，-1：未点赞'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帖子/评论和点赞用户关系表';

-- ----------------------------
-- Records of post_user_relation
-- ----------------------------

-- ----------------------------
-- Table structure for `rank_list_history`
-- ----------------------------
DROP TABLE IF EXISTS `rank_list_history`;
CREATE TABLE `rank_list_history` (
  `list_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一自增',
  `list_name` varchar(255) NOT NULL COMMENT '榜单名',
  `star_id` int(11) NOT NULL COMMENT '明星id，引用star_id',
  `date` date NOT NULL COMMENT '榜单日期',
  `rank` int(10) unsigned zerofill NOT NULL COMMENT '榜单排名',
  PRIMARY KEY (`list_id`),
  KEY `list_star_id` (`star_id`),
  CONSTRAINT `list_star_id` FOREIGN KEY (`star_id`) REFERENCES `star_info` (`star_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rank_list_history
-- ----------------------------

-- ----------------------------
-- Table structure for `star_info`
-- ----------------------------
DROP TABLE IF EXISTS `star_info`;
CREATE TABLE `star_info` (
  `star_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一自增',
  `star_name` varchar(255) NOT NULL COMMENT '明星的姓名',
  `img` varchar(255) DEFAULT NULL COMMENT '明星头像',
  `average_rank` int(10) unsigned zerofill NOT NULL DEFAULT '0000000000' COMMENT '今日平均排名，可以通过计算获得',
  `average_highest_rank` int(10) unsigned zerofill NOT NULL DEFAULT '0000000000' COMMENT '历史最高排名',
  `baidu_index` int(10) unsigned zerofill DEFAULT '0000000000',
  `current_weibofans_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT '微博今日粉丝数',
  `yesterday_weibofans_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT '微博昨天粉丝数',
  `current_insfans_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT 'ins今日粉丝数',
  `yesterday_insfans_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT 'ins昨天粉丝数',
  `tvshow_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT '明星参加综艺节目数',
  `ads_num` int(10) unsigned zerofill DEFAULT '0000000000' COMMENT '明星在接广告数',
  `banner` varchar(255) NOT NULL COMMENT '明星banner图',
  `identify` varchar(255) NOT NULL COMMENT '明星身份：歌手，演员等',
  `mv_num` int(11) DEFAULT '0' COMMENT '明星已拍摄mv数',
  PRIMARY KEY (`star_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='明星详细信息表';

-- ----------------------------
-- Records of star_info
-- ----------------------------
INSERT INTO `star_info` VALUES ('0', '易烊千玺', null, '0000000008', '0000000008', '0000009099', '0000010022', '0000023333', '0000032222', '0000009897', '0000000030', '0000000032', '', '', '0');
INSERT INTO `star_info` VALUES ('1', '吴亦凡', null, '0000000600', '0000000700', '0000000800', '0000000900', '0000001000', '0000000090', '0000000009', '0000000020', '0000000012', '', '', '0');
INSERT INTO `star_info` VALUES ('2', '蔡徐坤', null, '0000000008', '0000000009', '0000006756', '0000022133', '0000045333', '0000023456', '0000034523', '0000000021', '0000000010', '', '', '0');

-- ----------------------------
-- Table structure for `state`
-- ----------------------------
DROP TABLE IF EXISTS `state`;
CREATE TABLE `state` (
  `state_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '状态id，唯一自增',
  `account_name` varchar(255) NOT NULL COMMENT '发状态的账户名称',
  `content` varchar(255) DEFAULT NULL COMMENT '状态内容',
  `create_time` datetime NOT NULL COMMENT '状态创建时间',
  `imgs` varchar(255) DEFAULT NULL COMMENT '状态图片保存路径',
  `source` varchar(255) DEFAULT NULL COMMENT '状态来源',
  PRIMARY KEY (`state_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='明星状态表';

-- ----------------------------
-- Records of state
-- ----------------------------

-- ----------------------------
-- Table structure for `user_info`
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id，自增，唯一主键',
  `user_name` varchar(255) DEFAULT NULL COMMENT '用户名，用户的唯一标识符，不允许重复',
  `password` varchar(255) DEFAULT NULL COMMENT '用户密码',
  `img` varchar(255) DEFAULT NULL COMMENT '用户头像，如果用户没有上传则给定默认的',
  `suv` varchar(255) DEFAULT NULL COMMENT '用户suv，与user_id一一对应',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Records of user_info
-- ----------------------------

-- ----------------------------
-- Table structure for `user_list_relation`
-- ----------------------------
DROP TABLE IF EXISTS `user_list_relation`;
CREATE TABLE `user_list_relation` (
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `list_id` int(11) NOT NULL COMMENT '榜单id',
  `date` date NOT NULL COMMENT '打榜日期',
  `is_like` int(11) NOT NULL COMMENT '是否给该榜单打榜，0：打榜，1：未打榜',
  `star_id` int(11) NOT NULL COMMENT '明星id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户与榜单打榜的关系表';

-- ----------------------------
-- Records of user_list_relation
-- ----------------------------

-- ----------------------------
-- Table structure for `user_star_relation`
-- ----------------------------
DROP TABLE IF EXISTS `user_star_relation`;
CREATE TABLE `user_star_relation` (
  `user_id` int(11) NOT NULL COMMENT '用户id，引用用户表的id',
  `star_id` int(11) NOT NULL COMMENT '明星id，引用明星表的id',
  `follow_time` date DEFAULT NULL COMMENT '开始关注明星的时间',
  `support_num` int(11) DEFAULT NULL COMMENT '为明星应援次数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和明星的关系表';

-- ----------------------------
-- Records of user_star_relation
-- ----------------------------
