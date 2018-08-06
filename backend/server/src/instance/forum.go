package instance

import "time"

type post struct {
	post_id     int       `帖子id`
	user_id     int       `用户id `
	create_time time.Time `创建时间 `
	title       string    `标题 `
	content     string    `内容 `
	like_num    int       `点赞数 `
}

type connent struct {
	comment_id			int			`评论id `
	user_id				int			`用户id`
	parent_comment_id	int			`父评论id `
	create_time			time.Time	`时间 `
	content				string		`内容`
	like_num			int			`点赞数 `
	level				int			`级别 `

}
