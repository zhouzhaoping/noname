package instance

import "time"

type list struct {
	list_id		int			`id`
	list_name	string		`榜单名`
	star_id		int			`明星id `
	update_date	time.Time	`日期`
	rank		int			`排名`
}

type agenda struct {
	agenda_id   int       `行程id`
	star_id     int       `明星id `
	detail_time time.Time `时间`
	location    string    `地点 `
	content     string    `内容`
}

type production struct {
	pro_id   int    `商品id`
	pro_name string `商品名 `
	pro_type int    `代言商品还是同款商品`
}