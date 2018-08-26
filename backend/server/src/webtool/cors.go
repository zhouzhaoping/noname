package webtool

import "github.com/kataras/iris"

func SetCros(ctx iris.Context){
	ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT")
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
	ctx.ResponseWriter().Header().Set("content-type","application/json")//返回数据格式是json
}
