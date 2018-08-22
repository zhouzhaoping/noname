package user

import (
	"time"
	"sqltool"
	"github.com/kataras/iris"
)

type user_info struct {
	User_id		int		`xorm:"not null pk autoincr INT(11)"`
	User_name	string	`xorm:"not null VARCHAR(255)"`
	Password	string	`xorm:"not null VARCHAR(255)"`
	Img			string	`xorm:"not null VARCHAR(255)"`
	Suv			string	`xorm:"not null VARCHAR(255)"`
}

type login_log struct {
	Suv			string
	Login_time	time.Time
	Ip 			string
}

func NewUser_info(ctx iris.Context) *user_info {
	user := &user_info{
		User_name:ctx.FormValue("user_name"),
		Password:ctx.FormValue("password"),
		Img:ctx.FormValue("password"),
		Suv:ctx.GetCookie("SUV"),
	}
	id := ctx.PostValueIntDefault("user_id",-1)
	if id != -1{
		user.User_id = id
	}
	return user
}
func (user *user_info) checkPassword() (bool, *user_info) {
	user_find := new(user_info)
	if user.User_id > 0 {
		affected, err := sqltool.StarsuckEngine.ID(user.User_id).Get(user_find)
		if affected && err == nil && user_find.Password == user.Password {
			return true, user_find
		}
	} else if user.User_name != "" {
		affected, err := sqltool.StarsuckEngine.Where("user_name=?",user.User_name).Get(user_find)
		if affected && err == nil && user_find.Password == user.Password {
			return true, user_find
		}
	} else if user.Suv != "" {
		affected, err := sqltool.StarsuckEngine.Where("suv=?",user.Suv).Get(user_find)
		if affected && err == nil && user_find.Password == user.Password {
			return true, user_find
		} else {
			// new anonymous user signup
			user_find.Suv = user.Suv
			sqltool.StarsuckEngine.Insert(user_find)
			return true, user_find
		}
	}
	return false, user_find
}

type user_star_relation struct {
	user_id		int			`用户id`
	star_id		int			`明星id`
	follow_time	time.Time	`关注时间`
	support_num int			`应援次数`
}
