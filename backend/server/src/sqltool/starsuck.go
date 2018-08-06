package sqltool

import (
"fmt"
"github.com/go-xorm/xorm"
"github.com/go-xorm/core"
)

var starsuckEngine *xorm.Engine

func StarsuckInit() (debugOutput string){

	var err error

	//创建orm引擎
	starsuckEngine, err = xorm.NewEngine("mysql", "test:Starsuck8!@tcp(47.95.7.10:3306)/pickme_test?charset=utf8")

	//fmt.Println("engine:", engine)

	if err!=nil{
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}
	//连接测试
	if err := starsuckEngine.Ping(); err!=nil{
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	//日志打印SQL
	starsuckEngine.ShowSQL(true)

	//设置连接池的空闲数大小
	starsuckEngine.SetMaxIdleConns(5)
	//设置最大打开连接数
	starsuckEngine.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	starsuckEngine.SetTableMapper(core.SnakeMapper{})

	return
}

func StarsuckEnd(){
	starsuckEngine.Close()
}

