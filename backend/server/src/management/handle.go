package management

import (
	"github.com/kataras/iris"
	"fmt"
	"sqltool"
)

func GetUV(ctx iris.Context) {
	day_uv_find := make([]day_uv,0)
	err := sqltool.StarsuckEngine.Select("Date_FORMAT(login_time,\"%Y-%m-%d\"),count(distinct suv)").Table("login_log").
		Where("login_time >?","2018-08-30 10:00:00").GroupBy("DATE_FORMAT(login_time,\"%Y-%m-%d\")").Find(&day_uv_find)
	fmt.Println(err,day_uv_find)
	ctx.JSON(iris.Map{
		"state": "success",
		"data" : day_uv_find,
	})
	return
}

func GetIpCount(ctx iris.Context) {


	return
}