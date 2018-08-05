package sqltool

import (
	"database/sql"
	"fmt"
	"strconv"
)

func SqlTest() (debugOutput string) {

	debugOutput = "\n"

	db, err := sql.Open("mysql", "test:Starsuck8!@tcp(47.95.7.10:3306)/pickme_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	defer db.Close()

	var result sql.Result
	result, err = db.Exec("insert into t_test(age, name) values(?,?)", 13, "joe")
	if err != nil {
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("新插入记录的ID为", lastId)
	debugOutput += "\n" + "新插入记录的ID为" + strconv.Itoa(int(lastId))

	var row *sql.Row
	row = db.QueryRow("select * from t_test")
	var name string
	var id, age int
	err = row.Scan(&id, &age, &name)
	if err != nil {
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	fmt.Println(id, "\t", name, "\t", age)
	debugOutput += "\n" + strconv.Itoa(id) + "\t" + name + "\t" + strconv.Itoa(age)

	result, err = db.Exec("insert into t_test(age, name) values(?,?)", 24, "black")

	var rows *sql.Rows
	rows, err = db.Query("select * from t_test")
	if err != nil {
		fmt.Println(err)
		debugOutput += "\n" + err.Error()
		return
	}

	for rows.Next() {
		var name string
		var id, age int
		rows.Scan(&id, &age, &name)
		fmt.Println(id, "\t", name, "\t", age)
		debugOutput += "\n" + strconv.Itoa(id) + "\t" + name + "\t" + strconv.Itoa(age)
	}
	rows.Close()

	db.Exec("truncate table t_test")

	return
}
