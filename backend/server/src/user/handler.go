package user

import (
	"github.com/kataras/iris"
	"fmt"
	"sqltool"
	"time"
	"star"
	"strconv"
)


func PostUser(ctx iris.Context) {

	user := NewUser_info(ctx)
	fmt.Println("in postuser",user)
	user_find := new(User_info)

	if user.User_name == "" {
		ctx.JSON(iris.Map{
			"state":  "缺少名字",
		})
		return
	}

	// 检查名字
	yes, err := sqltool.StarsuckEngine.Where("user_name=?",user.User_name).Get(user_find)
	fmt.Println(user_find)
	if yes {
		ctx.JSON(iris.Map{
			"state":  "名字重复",
		})
		return
	}

	// 匿名用户更新
	yes, err = sqltool.StarsuckEngine.Where("suv=?",user.Suv).Get(user_find)
	fmt.Println(user_find)
	if yes && user_find.User_name == "" {
		user_find.User_name = user.User_name
		user_find.Password = user.Password
		user_find.Img = user.Img
		sqltool.StarsuckEngine.ID(user_find.User_id).Update(user_find)

		ctx.JSON(iris.Map{
			"state":  "success",
			"data": iris.Map{
				"user_id": user_find.User_id,
			},
		})
		return
	}

	// 新用户注册
	affected, err := sqltool.StarsuckEngine.Insert(user)
	fmt.Println(user)
	fmt.Println(affected, err)

	ctx.JSON(iris.Map{
		"state":  "success",
		"data": iris.Map{
			"user_id": user.User_id,
		},
	})
}

func PostLogin(ctx iris.Context){

	user := NewUser_info(ctx)
	fmt.Println(user)
	yes, user_find := user.checkPassword()
	fmt.Println(user_find)
	if yes {
		// update suv
		if user.Suv != "" && user.Suv != user_find.Suv {
			user_find.Suv = user.Suv
			sqltool.StarsuckEngine.ID(user_find.User_id).Update(user_find)
		}

		// log
		if user.Suv != ""{
			thislog := &login_log{
				user_find.Suv,
				time.Now(),
				ctx.RemoteAddr(),
				ctx.Request().Header.Get("origin"),
			}
			sqltool.StarsuckEngine.Insert(thislog)
		}

		ctx.JSON(iris.Map{
			"state":  "success",
			"data": iris.Map{
				"user_id": user_find.User_id,
			},
		})
	} else {
		ctx.JSON(iris.Map{
			"state":  "用户名或密码错误",
		})
	}
}

func GetUser(ctx iris.Context) {

	id,_ := ctx.Params().GetInt("user_id")
	user_find := new(User_info)

	yes, err := sqltool.StarsuckEngine.ID(id).Get(user_find)
	//fmt.Println(user_find)
	if yes && err == nil {
		//bytes, err := json.Marshal(user_find)
		//fmt.Println(err, string(bytes))

		if err == nil {
			//ctx.ContentType("application/json; charset=UTF-8")
			//ctx.Write(bytes)

			ctx.JSON(iris.Map{
				"state":  "success",
				"data": user_find,
			})
			return
		}
	}


	ctx.JSON(iris.Map{
		"state": "查无此人",
	})
}

func PutUser(ctx iris.Context) {

	user := NewUser_info(ctx)
	user.User_id, _ = ctx.Params().GetInt("user_id")
	fmt.Println(user)

	user_find := new(User_info)
	sqltool.StarsuckEngine.ID(user.User_id).Get(user_find)

	if user.User_name != "" && user.User_name !=user_find.User_name {
		user_find.User_name = user.User_name
	}
	if user.Password != "" && user.Password != user_find.Password {
		user_find.Password = user.Password
	}
	if user.Img != "" && user.Img != user_find.Img {
		user_find.Img = user.Img
	}
	if user.Suv != "" && user.Suv != user_find.Suv {
		user_find.Suv = user.Suv
	}
	sqltool.StarsuckEngine.ID(user.User_id).Update(user_find)

	ctx.JSON(iris.Map{
		"state":  "success",
	})
}

func GetFollowing(ctx iris.Context) {

	id,_ := ctx.Params().GetInt("user_id")

	stars := make([]star.Star_info_simple,0)
	err := sqltool.StarsuckEngine.Table("star_info").
									Join("INNER", "user_star_relation","star_info.star_id = user_star_relation.star_id").
									Where("user_star_relation.user_id=?",id).Asc("follow_time").
									Cols("star_info.star_id","star_name","img").Find(&stars)

	//ministars := make([]star.Star_info_simple,0)
	//ministars = stars

	if err == nil {
		fmt.Println(stars)
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  stars,
		})
	} else {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	}
}

func PutFollowing(ctx iris.Context) {

	user_id,_ := ctx.Params().GetInt("user_id")
	star_id,err := strconv.Atoi(ctx.FormValue("star_id"))

	fmt.Println(err, star_id)


	u_s := new(user_star_relation)
	yes, err := sqltool.StarsuckEngine.Where("user_id=? and star_id=?",user_id,star_id).Get(u_s)
	fmt.Println(yes,err)
	if yes {
		ctx.JSON(iris.Map{
			"state":  "不可重复关注",
		})
		return
	}

	u_s = &user_star_relation{
		user_id,
		star_id,
		time.Now(),
		0,
	}
	sqltool.StarsuckEngine.Insert(u_s)
	ctx.JSON(iris.Map{
		"state":  "success",
	})
}

func PutUnFollowing(ctx iris.Context) {

	user_id,_ := ctx.Params().GetInt("user_id")
	star_id,_ := strconv.Atoi(ctx.FormValue("star_id"))


	u_s := new(user_star_relation)
	yes, _ := sqltool.StarsuckEngine.Where("user_id=? and star_id=?",user_id,star_id).Get(u_s)
	if !yes {
		ctx.JSON(iris.Map{
			"state":  "您还未关注",
		})
		return
	}

	sqltool.StarsuckEngine.Delete(u_s)
	ctx.JSON(iris.Map{
		"state":  "success",
	})
}

func GetIsAnonymous(ctx iris.Context) {
	user_id, _ := ctx.Params().GetInt("user_id")
	user_find := new(User_info)
	yes, err := sqltool.StarsuckEngine.ID(user_id).Get(user_find)

	if err != nil{
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
		return
	}
	if !yes {
		ctx.JSON(iris.Map{
			"state": "查无此人",
		})
		return
	}else {
		ctx.JSON(iris.Map{
			"state":  "success",
			"is_anonymous":user_find.User_name == "",
		})
	}
}
