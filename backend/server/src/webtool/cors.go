package webtool

import (
	"github.com/kataras/iris"
	"strings"
)

func SetCros(ctx iris.Context){
	Origin := ctx.Request().Header.Get("Origin")
	if 0 != len(Origin) && strings.Contains(Origin, "domain.com"){
		ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", Origin)
		ctx.ResponseWriter().Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		ctx.ResponseWriter().Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type")
		ctx.ResponseWriter().Header().Add("Access-Control-Allow-Credentials", "true")
		ctx.ResponseWriter().Header().Add("content-type","application/json")//返回数据格式是json
	}
	//ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
	//ctx.ResponseWriter().Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT")
	//ctx.ResponseWriter().Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
	//ctx.ResponseWriter().Header().Set("content-type","application/json")//返回数据格式是json
}
