package handler

import "github.com/kataras/iris"

func GetPosts(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")

	ctx.Writef("user_id: %s", user_id)

	hehe := iris.Map{
		"post_id":"postId",
		"create_time":"2018-03-20 23:11:23",
		"title":"title",
		"content":"XXX粉丝XXXX",
		"like_num":2000,
	}

	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : []iris.Map{hehe,hehe,hehe},
	})
}

func PostPosts(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")

	ctx.Writef("user_id: %s", user_id)

	hehe := iris.Map{
		"post_id":"postId",
		"create_time":"2018-03-20 23:11:23",
		"title":"title",
		"content":"XXX粉丝XXXX",
		"like_num":2000,
	}

	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : []iris.Map{hehe,hehe,hehe},
	})
}

func GetComments(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")
	post_id := ctx.URLParamDefault("post_id", "post_id")

	ctx.Writef("user_id: %s post_id: %s", user_id, post_id)

	hehe := iris.Map{
		"post_id":"postId",
		"create_time":"2018-03-20 23:11:23",
		"title":"title",
		"content":"XXX粉丝XXXX",
		"like_num":2000,
	}

	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : []iris.Map{hehe,hehe,hehe},
	})
}


