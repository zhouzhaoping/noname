package instance

import "time"

type state struct {
	state_id		int			`动态id`
	account_id		int			`账号id`
	account_name	string		`账户名称 `
	content			string		`内容`
	create_time		time.Time	`发表时间 `
	imgs			string		`图片id们`
	source			string		`来源`
}

type news struct {
	news_id		int		`资讯`
	star_id		int		`明星id `
	img			string	`图片id`
	title		string	`标题 `
	url			string	`链接`
	source		string	`来源 `
	create_time	string	`发表时间`
}
