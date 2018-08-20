package handler

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"fmt"
)

func News(ctx iris.Context) {
	//user_id := ctx.URLParamDefault("user_id", "anonymous")
	star_id := ctx.URLParamDefault("star_id",	 "follow")

	//ctx.Writef("user_id: %s star_id: %s", user_id, star_id)
	/*starnew  := iris.Map{
		"title":"title",
		"url":"url",
		"img":"img",
		"create_time":"create_time",
		"source":"source",
	}*/

	if star_id == "follow"{
		star_id = "0"
	}
	filename := "news_list" + star_id + ".json"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		ctx.Writef(err.Error())
	}
	ctx.Write(bytes)
	/*ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : bytes,
	})*/
}

func States(ctx iris.Context) {
	user_id := ctx.URLParamDefault("user_id", "anonymous")
	star_id := ctx.URLParamDefault("star_id",	 "0")

	ctx.Writef("user_id: %s star_id: %s", user_id, star_id)

	agenda := iris.Map{
		"agenda_id":"002",
		"star_id":"20002",
		"detailtime":"2018-08-29 12:30:33",
		"location":"北京南苑机场",
		"content":"XXX",
		"type":0,
	}
	state := iris.Map{
		"state_id":9100,
		"account_id":"504",
		"account_name":"吴亦凡官方微博",
		"create_time":"2018-03-20 23:11:23",
		"content":"XXX粉丝XXXX",
		"img":"img",
		"source":"微博",
	}

	ctx.JSON(iris.Map{
		"state" : 10000,
		"msg" : "success",
		"data" : iris.Map{
			"agenda":agenda,
			"staes":[]iris.Map{state,state,state},
		},
	})
}

