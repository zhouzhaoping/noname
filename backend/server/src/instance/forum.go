package instance

import "time"

/*
CREATE TABLE `post` (
`post_id` int(11) NOT NULL AUTO_INCREMENT, 	--帖子id
`user_id` int(10) NOT NULL,					--用户id
`create_time` datetime DEFAULT NULL,		--创建时间
`title` varchar(255) DEFAULT '',			--标题
`content` varchar(255) DEFAULT '',			--内容
`like_num` int(10) unsigned zerofill DEFAULT NULL,	--点赞数
PRIMARY KEY (`post_id`),
KEY `post_user_id` (`user_id`),
CONSTRAINT `post_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON UPDATE CASCADE
)
*/

type post struct {
	post_id     int			`xorm:"not null pk autoincr INT(11)"`
	user_id     int			`xorm:"not null INT(10)"`
	create_time time.Time	`xorm:"DATE"`
	title       string
	content     string
	like_num    int
}

type comment struct {
	comment_id			int			`评论id `
	user_id				int			`用户id`
	post_id				int
	parent_comment_id	int			`父评论id `
	create_time			time.Time	`时间 `
	content				string		`内容`
	like_num			int			`点赞数 `
	level				int			`级别 `

}
