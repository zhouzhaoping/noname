package handler

import (
	"github.com/kataras/iris"
)

func Stars(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")
	follow := ctx.URLParamDefault("follow",	 "all")

	ctx.Writef("user_id: %s follow: %s", user_id, follow)
	hehe := iris.Map{
		"star_id":"starId",
		"name":"name",
		"follow_time":"follow_time",
	}

	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : []iris.Map{hehe,hehe,hehe},
	})
}



// 明星个人动态页

// 明星打榜页

// 发帖接口


// 评论和回复评论接口