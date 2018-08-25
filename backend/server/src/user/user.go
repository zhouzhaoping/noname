package user

import (
	"time"
	"sqltool"
	"github.com/kataras/iris"
)

type user_info struct {
	User_id		int		`xorm:"not null pk autoincr INT(11)" json:"user_id"`
	User_name	string	`xorm:"default null VARCHAR(255)" json:"user_name"`
	Password	string	`xorm:"default null VARCHAR(255)" json:"password"`
	Img			string	`xorm:"default null VARCHAR(255)" json:"img"`
	Suv			string	`xorm:"default null VARCHAR(255)" json:"suv"`
}

type login_log struct {
	Suv			string		`xorm:"not null VARCHAR(255)" json:"suv"`
	Login_time	time.Time	`xorm:"not null DATETIME" json:"login_time"`
	Ip 			string		`xorm:"default null VARCHAR(255)" json:"ip"`
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
			// 新增匿名用户
			user_find.Suv = user.Suv
			sqltool.StarsuckEngine.Insert(user_find)
			return true, user_find
		}
	}
	return false, user_find
}

type user_star_relation struct {
	User_id		int			`xorm:"not null INT(11)" json:"user_id"`
	Star_id		int			`xorm:"not null INT(11)" json:"star_id"`
	Follow_time	time.Time	`xorm:"default null DATE" json:"follow_time"`
	Support_num int			`xorm:"default null INT(11)" json:"support_num"`
}

type auth_accounts struct {
	User_id      int    	`xorm:"not null INT(11)" json:"user_id"`
	User_name    string 	`json:"user_name"`
	Password     string 	`json:"password"`
	Account_type int    	`json:"account_type"` //0：百度，1：微博，2：ins'
}

type user_list_relation struct {
	User_id int       `xorm:"not null INT(11)" json:"user_id"`
	List_id int       `json:"list_id"`
	Date    time.Time `json:"date"`
	Is_like int       `json:"is_like"` //0：打榜，1：未打榜',
	Star_id int       `json:"star_id"`
}
