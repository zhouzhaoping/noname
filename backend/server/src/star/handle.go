package star

import (
	"github.com/kataras/iris"
	"fmt"
	"sqltool"
)

func GetStars(ctx iris.Context) {
	id,_ := ctx.Params().GetInt("user_id")
	stars := make([]Star_info_simple,0)
	err := sqltool.StarsuckEngine.Table("star_info").
		Cols("star_id", "star_name","img").Asc("star_name").Find(&stars)

	for i,v := range(stars){
		stars[i].Is_Following, err = sqltool.StarsuckEngine.Table("user_star_relation").Where("user_id = ? and star_id=?",id,v.Star_id).Exist()
	}

	if err == nil{
		fmt.Println(stars)
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  stars,
		})
	}else{
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "数据库错误",
		})
	}
}

func GetStar(ctx iris.Context) {
	id,_ := ctx.Params().GetInt("star_id")
	fmt.Println(id)
	star_find := new(Star_info)
	yes, err := sqltool.StarsuckEngine.ID(id).Get(star_find)
	if yes&& err==nil{
		ctx.JSON(iris.Map{
			"state": "success",
			"data":  star_find,
		})
	}else {
		fmt.Println(err)
		ctx.JSON(iris.Map{
			"state": "查无此人",
		})
	}
}
