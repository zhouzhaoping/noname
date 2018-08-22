package user

import (
	"github.com/kataras/iris"
	"fmt"
	"sqltool"
	"time"
	"encoding/json"
)


func PostUser(ctx iris.Context) {
	user := NewUser_info(ctx)
	fmt.Println("in postuser",user)
	user_find := new(user_info)

	if user.User_name == "" {
		ctx.JSON(iris.Map{
			"state":  "缺少名字",
		})
		return
	}

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
			"user_id": user_find.User_id,
		})
		return
	}

	// 新用户注册
	affected, err := sqltool.StarsuckEngine.Insert(user)
	fmt.Println(user)
	fmt.Println(affected, err)

	ctx.JSON(iris.Map{
		"user_id":  user.User_id,
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
			}
			sqltool.StarsuckEngine.Insert(thislog)
		}

		ctx.JSON(iris.Map{
			"user_id":  user_find.User_id,
		})
	} else {
		ctx.JSON(iris.Map{
			"state":  "用户名或密码错误",
		})
	}
}

func GetUser(ctx iris.Context) {
	id,_ := ctx.Params().GetInt("user_id")
	user_find := new(user_info)

	yes, err := sqltool.StarsuckEngine.ID(id).Get(user_find)
	//fmt.Println(user_find)
	if yes && err == nil {
		bytes, err := json.Marshal(user_find)
		fmt.Println(err, string(bytes))

		if err == nil {
			ctx.ContentType("application/json; charset=UTF-8")
			ctx.Write(bytes)
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

	user_find := new(user_info)
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