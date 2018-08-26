package forum

import (
	"time"
	"github.com/kataras/iris"
)

type Post struct {
	Post_id				int			`xorm:"not null pk autoincr INT(11)" json:"post_id"`
	User_id				int 		`json:"user_id"`
	Create_time			time.Time	`json:"create_time"`
	Title				string 		`json:"title"`
	Content				string 		`json:"content"`
	Like_num			uint 		`json:"like_num"`
	Star_id				int			`json:"star_id"`
	Parent_comment_id	int			`json:"parent_comment_id"`
	Comment_num			int 		`json:"comment_num"`
	Level				int			`json:"level"`
	Imgs				string		`json:"imgs"`
}

func NewPost(ctx iris.Context) *Post {
	post := &Post{
		Title:ctx.FormValue("title"),
		Content:ctx.FormValue("content"),
		Imgs:ctx.GetCookie("imgs"),
		Like_num:0,
		Comment_num:0,
	}
	id := ctx.PostValueIntDefault("user_id",-1)
	if id != -1{
		post.User_id = id
	}
	id = ctx.PostValueIntDefault("star_id",-1)
	if id != -1{
		post.Star_id = id
	}
	return post
}

type post_like struct {
	post			Post		`json:"post"`
	is_like			bool		`json:"is_like"`
}

type post_user_relation struct {
	Post_id int `json:"post_id"`
	User_id int `json:"user_id"`
	Is_like int `json:"is_like"` //0：点赞，1：未点赞
}
