package handler

import (
	"github.com/kataras/iris"
	"fmt"
)

// 首页hot推荐:所有已关注明星的资讯，按照时间排序，50条
func Hot(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")
	fmt.Println(user_id)

	hehe := iris.Map{
		"news_id":"newId",
		"star_id":"starId",
		"title":"title",
		"url":"url",
		"source":"source",
	}
	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : []iris.Map{hehe,hehe,hehe},
	})
}

// 明星资讯页


// 饭圈首页

// 明星个人动态页

// 明星打榜页

// 发帖接口


// 评论和回复评论接口