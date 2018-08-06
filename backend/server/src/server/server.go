package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	_ "github.com/go-xorm/xorm"
	_ "github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"io/ioutil"
	"strconv"

	"sqltool"
	"imagetool"
	"handler"
)

func main() {
	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("server.pid", []byte(strconv.Itoa(pid)), 0777)
	}
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) //, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			case syscall.SIGHUP:
				fmt.Println("sighup", s)
				//case syscall.SIGUSR1:
				//    fmt.Println("usr1", s)
				//case syscall.SIGUSR2:
				//    fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")

	// 创建图片存储参数
	imagetool.LoadConf()
	//创建orm引擎
	//sqltool.XormInit()
	sqltool.StarsuckInit()

	go func() {
		sum := 0
		for {
			sum++
			//fmt.Println("sum:", sum)
			time.Sleep(time.Second)
		}
	}()

	var app *iris.Application = iris.New()

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	ServerTestBinder(app)
	Binder(app)

	// 图片服务器
	// Method:   GET
	// Resource: http://localhost:8080/testimage
	app.Get("/testimage", func(ctx iris.Context) {
		imagetool.HomeHandler(ctx.ResponseWriter(), ctx.Request())
	})
	// Method:   POST
	// Resource: http://localhost:8080/testimage
	app.Post("/testimage", func(ctx iris.Context) {
		imagetool.UploadHandler(ctx.ResponseWriter(), ctx.Request())
	})
	// Method:   GET
	// Resource: http://localhost:8080/testimage/{imgid}
	app.Get("/testimage/{imgid:string}", func(ctx iris.Context) {
		imgid := ctx.Params().Get("imgid")
		imagetool.DownloadHandler(ctx.ResponseWriter(), ctx.Request(), imgid)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	//<-sigTERM
	//log.Print("killed")
}

func Binder(app *iris.Application){
	// Method:   GET
	// Resource: http://localhost:8080/hot?user_id=anonymous
	app.Handle("GET", "/hot", handler.Hot)
}
func ServerTestBinder(app *iris.Application){
	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/testsql
	app.Get("/testsql", func(ctx iris.Context) {
		//ctx.WriteString(sqltool.SqlTest())
		ctx.WriteString(sqltool.XormTest())
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// Method:   GET
	// Resource: http://localhost:8080/user/john not /user/ or /user
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	// Method:   POST
	// Resource: http://localhost:8080/user/john/ and /user/john/send
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	// Method:   GET(query)
	// Resource: http://localhost:8080/welcome?firstname=Jane&lastname=Doe.
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		// shortcut for ctx.Request().URL.Query().Get("lastname").
		lastname := ctx.URLParam("lastname")//可以为空

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	// Method:   POST(post form)
	// Resource: http://localhost:8080/form_post
	// message=hello&nick=john
	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// Method:   POST(query and post form)
	// Resource: http://localhost:8080/post?id=1234&page=1
	// name=manu&message=this_is_great // body 里的参数
	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		// or `ctx.PostValue` for POST, PUT & PATCH-only HTTP Methods.

		app.Logger().Infof("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
}

func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")

	os.Remove("server.pid")
	//销毁orm引擎
	//sqltool.XormEnd()
	sqltool.StarsuckEnd()

	fmt.Println("结束退出...")
	os.Exit(0)
}

