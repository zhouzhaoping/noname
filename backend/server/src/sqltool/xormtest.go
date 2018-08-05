package sqltool

import (
	"time"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"strconv"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Address  string    `xorm:"VARCHAR(256)"`
}

var engine *xorm.Engine

func XormTest() (debugOutput string){

	//fmt.Println("engine:", engine)

	debugOutput = "\n"

	fmt.Println("insert begin")
	//增
	user := new(User)
	user.Username="tyming1"
	affected, err := engine.Insert(user)
	fmt.Println(affected)
	debugOutput += "\n" + strconv.Itoa(int(affected))
	fmt.Println("insert end")

	fmt.Println("insert begin")
	//增
	user = new(User)
	user.Username="tyming2"
	affected, err = engine.Insert(user)
	fmt.Println(affected)
	debugOutput += "\n" + strconv.Itoa(int(affected))
	fmt.Println("insert end")

	fmt.Println("delete begin")
	//删
	user = new(User)
	user.Username="tyming1"
	affected_delete,err := engine.Delete(user)
	fmt.Println(affected_delete)
	debugOutput += "\n" + strconv.Itoa(int(affected_delete))
	fmt.Println("delete end")

	fmt.Println("change begin")
	//改
	user = new(User)
	user.Username="tyming2"
	affected_update,err := engine.Id(1).Update(user)// id 1 乐观锁
	fmt.Println(affected_update)
	debugOutput += "\n" + strconv.Itoa(int(affected_update))
	fmt.Println("change end")

	fmt.Println("find begin")
	//查
	user = new(User)
	//result,err := engine.Id(1).Get(user)
	result,err := engine.Where("Username=?","tyming2").Get(user)
	fmt.Println(result)
	debugOutput += "\n" + strconv.FormatBool(result)
	fmt.Println("find end")

	fmt.Println(err)
	return
}

func XormInit() (debugOutput string){

	var err error

	//创建orm引擎
	engine, err = xorm.NewEngine("mysql", "test:Starsuck8!@tcp(47.95.7.10:3306)/pickme_test?charset=utf8")

	//fmt.Println("engine:", engine)

	if err!=nil{
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}
	//连接测试
	if err := engine.Ping(); err!=nil{
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})

	return
}

func XormEnd(){
	engine.Close()
}
