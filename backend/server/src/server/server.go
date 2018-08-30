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
	"io/ioutil"
	"strconv"

	"sqltool"
	"imagetool"
	"handler"
	"user"
	"star"
	"news_states"
	"forum"
	"webtool"
	"time"
	"github.com/kataras/iris/cache"
	"management"
	"updater"
)

func main() {
	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("server.pid", []byte(strconv.Itoa(pid)), 0777)
	}
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
 	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			case syscall.SIGHUP:
				fmt.Println("sighup", s)
			case syscall.SIGUSR1:
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")

	// 创建图片存储参数
	imagetool.LoadConf()
	// 缓存初始化
	imagetool.CacheInit()
	// 创建orm引擎
	sqltool.StarsuckInit()

	go func() {
		//sum := 0
		for {
			//sum++
			// fmt.Println("sum:", sum)
			//handler.Refresh(0)
			//handler.UpdateNews()


			fmt.Println("update...")
			updater.NewsUpdater()
			updater.StatesUpdater()
			time.Sleep(time.Hour * 1)
		}
	}()

	var app *iris.Application = iris.New()

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	//app.Use(cache.Handler(2*time.Second))
	//ServerTestBinder(app)
	//Binder(app)
	CoreBinder(app)

	// 图片服务器
	// Method:   GET
	// Resource: http://localhost:8080/image
	//app.Get("/image", func(ctx iris.Context) {
	//	imagetool.HomeHandler(ctx.ResponseWriter(), ctx.Request())
	//})

	// Method:   POST
	// Resource: http://localhost:8080/image
	app.Post("/api/image", imagetool.UploadHandler)

	// Method:   GET
	// Resource: http://localhost:8080/image/{imgid}
	app.Get("/api/image/{imgid:string}", cache.Handler(3*time.Second),imagetool.DownloadHandler)

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	//<-sigTERM
	//log.Print("killed")

}

func CoreBinder(app *iris.Application){


	// page
	app.Handle("GET","/*",webtool.GetPage)
	app.StaticWeb("/static", "/root/pickme/frontend/build/static")

	// user system
	app.Handle("POST","/api/user",user.PostUser)
	app.Handle("POST","/api/login",user.PostLogin)
	app.Handle("GET","/api/user/{user_id:int}",user.GetUser)
	app.Handle("PUT","/api/user/{user_id:int}",user.PutUser)
	app.Handle("GET","/api/user/{user_id:int}/isanonymous",user.GetIsAnonymous)

	// user_star
	app.Handle("GET","/api/user/{user_id:int}/following",user.GetFollowing)
	app.Handle("PUT","/api/user/{user_id:int}/following",user.PutFollowing)
	app.Handle("PUT","/api/user/{user_id:int}/unfollowing",user.PutUnFollowing)

	// star system
	app.Handle("GET","/api/star/user/{user_id:int}",star.GetStars)
	app.Handle("GET","/api/star/{star_id:int}",star.GetStar)

	// news and states
	app.Handle("GET","/api/user/{user_id:int}/news",news_states.GetUserNews)
	app.Handle("GET","/api/star/{star_id:int}/news",news_states.GetStarNews)
	app.Handle("GET","/api/star/{star_id:int}/states",news_states.GetStarStates)



	// forum
	app.Handle("GET","/api/star/{star_id:int}/head",forum.GetStarHead)
	app.Handle("GET","/api/star/{star_id:int}/posts/user/{user_id:int}",forum.GetStarPost)
	app.Handle("GET","/api/post/{post_id:int}/user/{user_id:int}",forum.GetPost)
	app.Handle("POST","/api/post",forum.PostNewPost)
	app.Handle("POST","/api/post/{post_id:int}",forum.PostReplyPost)
	app.Handle("PUT","/api/post/{post_id:int}/like",forum.PutPostLike)
	app.Handle("PUT","/api/post/{post_id:int}/unlike",forum.PutPostUnLike)


	// TODO
	app.Handle("GET","/api/management/uv",management.GetUV)
	app.Handle("GET","/api/management/ipcount",management.GetIpCount)

	// Method:   GET
	// Resource: http://localhost:8080/news?user_id=anonymous&star_id=follow
	// 首页hot推荐:所有已关注明星的资讯，按照时间排序，50条
	// http://localhost:8080/news?user_id=anonymous&star_id=follow
	// 某个明星资讯页:回目标明星相关的50条资讯，按照时间排序
	// http://localhost:8080/news?user_id=anonymous&star_id=0
	//app.Handle("GET", "/news", handler.News)

	// Method:   GET
	// Resource: http://localhost:8080/states?user_id=anonymous&star_id=0
	// 明星个人动态页
	// 1、显示该明星最近的行程信息
	// 2、显示该明星相关的账号所发的状态
	//app.Handle("GET", "/states", handler.States)
}
func Binder(app *iris.Application){
	// Method:   GET
	// Resource: http://localhost:8080/user_id=anonymous&stars?follow=all
	// 返回所有关注的明星的名字，头像
	// follow = my 为已关注的，按照关注时间排序；follow = all 为所有的
	app.Handle("GET", "/stars", handler.Stars)

	// Method:   GET
	// Resource: http://localhost:8080/news?user_id=anonymous&star_id=follow
	// 首页hot推荐:所有已关注明星的资讯，按照时间排序，50条
	// 某个明星资讯页:回目标明星相关的50条资讯，按照时间排序
	// star_id = follow 则为已关注
	app.Handle("GET", "/news", handler.News)

	// Method:   GET
	// Resource: http://localhost:8080/forum/getposts?user_id=anonymous
	// 饭圈首页:返回当前最新的50条帖子，按照时间顺序排列
	app.Handle("GET", "/forum/getposts", handler.GetPosts)

	// Method:   GET
	// Resource: http://localhost:8080/forum/getcomments?user_id=anonymous&post_id=post_id
	// 饭圈首页:返回当前帖子的评论以及子评论，时间排序
	app.Handle("GET", "/forum/getcomments", handler.GetComments)

	// Method:   POST
	// Resource: http://localhost:8080/forum/postpost
	// user_id=anonymous
	// title=title
	// content=content
	// 发表一个帖子
	//app.Handle("GET", "/forum/postpost", handler.PostPost)


	// Method:   POST
	// Resource: http://localhost:8080/forum/postcomment
	// user_id=anonymous
	// title=title
	// content=content
	// 饭圈首页:返回当前帖子的评论以及子评论，时间排序
	//app.Handle("GET", "/forum/postpost", handler.PostPost)
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

