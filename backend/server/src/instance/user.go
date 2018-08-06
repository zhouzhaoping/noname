package instance

import "time"

type user_info struct {
	user_id		int		`xorm:"not null pk autoincr INT(11)"`
	user_name	string	`xorm:"not null VARCHAR(32)"`
	password	string	`xorm:"not null VARCHAR(32)"`
}

type user_star_relation struct {
	user_id		int			`用户id`
	star_id		int			`明星id`
	follow_time	time.Time	`关注时间`
	support_num int			`应援次数`
}
