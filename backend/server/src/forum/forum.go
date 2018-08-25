package forum

import "time"

type post struct {
	Post_id				int			`xorm:"not null pk autoincr INT(11)" json:"post_id"`
	User_id				int 		`json:"user_id"`
	Create_time			time.Time	`json:"create_time"`
	Title				string 		`json:"title"`
	Content				string 		`json:"content"`
	Like_num			int 		`json:"like_num"`
	Star_id				string		`json:"star_id"`
	Parent_comment_id	int			`json:"parent_comment_id"`
	Comment_num			int 		`json:"comment_num"`
	Level				int			`json:"level"`
	Imgs				string		`json:"imgs"`
}

type post_user_relation struct {
	Post_id int `json:"post_id"`
	User_id int `json:"user_id"`
	Is_like int `json:"is_like"` //0：点赞，1：未点赞
}
