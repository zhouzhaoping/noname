package webtool

import (
	"github.com/kataras/iris"
	"io/ioutil"
)

func GetPage(ctx iris.Context) {
	indexHTML,_ := ioutil.ReadFile("/root/pickme/frontend/build/index.html")
	ctx.HTML(string(indexHTML))
}
