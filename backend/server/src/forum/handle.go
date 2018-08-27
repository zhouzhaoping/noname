package forum

import (
	"github.com/kataras/iris"
	"star"
	"sqltool"
	"fmt"
)

func GetStarHead(ctx iris.Context) {

	id,_ := ctx.Params().GetInt("star_id")
	star_info_find := new(star.Star_info)

	yes, err := sqltool.StarsuckEngine.Table("star_info").ID(id).Cols("star_id","banner","identify").Get(star_info_find)
	if yes && err == nil{
		fmt.Println(star_info_find)
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  iris.Map{
				"star_id":star_info_find.Star_id,
				"banner":star_info_find.Banner,
				"identify":star_info_find.Identify,
			},

		})
	}else {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	}
}
func GetStarPost(ctx iris.Context) {
	star_id,_ := ctx.Params().GetInt("star_id")
	user_id,_ := ctx.Params().GetInt("user_id")
	post_find := make([]Post,0)


	err := sqltool.StarsuckEngine.Where("star_id",star_id).Desc("create_time").Find(&post_find)
	if err == nil {
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}

	post_like_find := make([]post_like,len(post_find))
	for i,v := range(post_find){
		post_like_find[i].post = v
		post_like_find[i].is_like, err = sqltool.StarsuckEngine.Table("post_user_relation").Where("user_id = ? and post_id=? and is_like=?",user_id,v.Post_id,0).Exist()

		if err == nil {
			ctx.JSON(iris.Map{
				"state": "数据库错误",
			})
			return
		}
	}

	fmt.Println(err)
	ctx.JSON(iris.Map{
		"state": "success",
		"data":post_like_find,
	})
}
func GetPost(ctx iris.Context) {
	post_id,_ := ctx.Params().GetInt("post_id")
	user_id,_ := ctx.Params().GetInt("user_id")

	post_like_find := new(post_like)
	comment_find := make([]Post,0)

	sqltool.StarsuckEngine.ID(post_id).Get(post_like_find.post)
	post_like_find.is_like, _ = sqltool.StarsuckEngine.Table("post_user_relation").Where("user_id = ? and post_id=? and is_like=?",user_id,post_id,0).Exist()
	sqltool.StarsuckEngine.Where("parent_comment_id=?",post_id).Find(&comment_find)

	comment_like_find := make([]post_like,len(comment_find))
	for i,v := range(comment_find) {
		var err error
		comment_like_find[i].post = v
		comment_like_find[i].is_like, err = sqltool.StarsuckEngine.Table("post_user_relation").Where("user_id = ? and post_id=? and is_like=?",user_id,v.Post_id,0).Exist()
		if err == nil {
			ctx.JSON(iris.Map{
				"state": "数据库错误",
			})
			return
		}
	}

	ctx.JSON(iris.Map{
		"state": "success",
		"data":iris.Map{
			"post_detail":post_like_find,
			"comment_detail":comment_like_find,
		},
	})
}
func PostNewPost(ctx iris.Context) {
	//post := NewPost(ctx)

}
func PostReplyPost(ctx iris.Context) {

}
func PutPostLike(ctx iris.Context) {
	post_id,_ := ctx.Params().GetInt("post_id")
	user_id := ctx.PostValueIntDefault("user_id",-1)

	p_u := new(post_user_relation)
	yes, err := sqltool.StarsuckEngine.Where("user_id = ? and post_id=?",user_id,post_id).Get(p_u)
	if err!= nil {
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}
	if yes {
		if p_u.Is_like == 0{
			ctx.JSON(iris.Map{
				"state": "不可重复点赞" ,
			})
			return
		} else {
			p_u.Is_like = 0
			_, err = sqltool.StarsuckEngine.Update(p_u)
			if err!= nil {
				ctx.JSON(iris.Map{
					"state": "数据库错误",
				})
				return
			}
			ctx.JSON(iris.Map{
				"state": "success" ,
			})
			return
		}
	} else {
		p_u.Post_id = post_id
		p_u.User_id = user_id
		p_u.Is_like = 0
		sqltool.StarsuckEngine.Insert(p_u)
		ctx.JSON(iris.Map{
			"state": "success" ,
		})
		return
	}
}
func PutPostUnLike(ctx iris.Context) {
	post_id,_ := ctx.Params().GetInt("post_id")
	user_id := ctx.PostValueIntDefault("user_id",-1)

	p_u := new(post_user_relation)
	yes, err := sqltool.StarsuckEngine.Where("user_id = ? and post_id=?",user_id,post_id).Get(p_u)
	if err!= nil {
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}
	if yes {
		if p_u.Is_like == 1{
			ctx.JSON(iris.Map{
				"state": "并没有点赞" ,
			})
			return
		} else {
			p_u.Is_like = 1
			_, err = sqltool.StarsuckEngine.Update(p_u)
			if err!= nil {
				ctx.JSON(iris.Map{
					"state": "数据库错误",
				})
				return
			}
			ctx.JSON(iris.Map{
				"state": "success" ,
			})
			return
		}
	} else {
		p_u.Post_id = post_id
		p_u.User_id = user_id
		p_u.Is_like = 1
		sqltool.StarsuckEngine.Insert(p_u)
		ctx.JSON(iris.Map{
			"state": "success" ,
		})
		return
	}
}