package news_states

import (
	"github.com/kataras/iris"
	"sqltool"
	"fmt"
)

func GetUserNews(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("user_id")
	var u_s []int
	err := sqltool.StarsuckEngine.Table("user_star_relation").Cols("star_id").Where("user_id=?", id).Find(&u_s)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}

	fmt.Println(u_s)
	news_find := make([]News, 0)
	if len(u_s) != 0 { //已经有关注的明星
		err = sqltool.StarsuckEngine.In("star_id",u_s).Desc("create_time").Limit(50).Find(&news_find)
	} else { //没有关注明星
		err = sqltool.StarsuckEngine.Desc("create_time").Limit(50).Find(&news_find)
	}

	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	} else {
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  news_find,
		})
	}
}

func GetStarNews(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("star_id")
	news_find := make([]News, 0)
	err := sqltool.StarsuckEngine.Where("star_id=?",id).Desc("create_time").Limit(50).Find(&news_find)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	} else {
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  news_find,
		})
	}
}

func GetStarStates(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("star_id")

	agenda_find := new(Agenda)
	_,err := sqltool.StarsuckEngine.Where("star_id=?",id).Desc("detail_time").Limit(1).Get(agenda_find)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}
	fmt.Println(agenda_find)


	var accountnames []string
	states_find := make([]State, 0)
	err = sqltool.StarsuckEngine.Table("info_source").Cols("account_name").Where("star_id=? and usage_type=? and source=?", id,1,"weibo").Find(&accountnames)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}
	fmt.Println(accountnames)
	err = sqltool.StarsuckEngine.In("account_name",accountnames).Desc("create_time").Limit(50).Find(&states_find)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	} else {
		ctx.JSON(iris.Map{
			"state": "success",
			"data": iris.Map{
				"agenda": agenda_find,
				"states": states_find,
			},
		})
	}
}