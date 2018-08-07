package handler

import (
	"time"
	"os/exec"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func Refresh(star_id int) {
	lasttime := findLastTime(star_id)
	if lasttime == nil{
		cmd := exec.Command("python", "../backend/informarion_spider/starspider.py")
		buf, err := cmd.Output()
		fmt.Printf("%s\n%s",buf,err)
	}else{

	}
}
func findLastTime(star_id int) (lasttime *time.Time){

	return nil
}
func UpdateNews(){
	fmt.Println("update...")
	starnewsArray, err := readFile("D:\\pickme\\backend\\informarion_spider\\news_list.json")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		return
	}
	fmt.Println(starnewsArray)
	//sqltool.StarsuckEngine
}

//结构体里的字段首字母必须大写，否则无法正常解析
type Starnew struct {
	Title string  			`json:"title"`
	Url string				`json:"url"`
	Img string				`json:"img"`
	Create_time string		`json:"create_time"`
	Source string			`json:"source"`
}

func readFile(filename string) ([]Starnew, error) {
	fmt.Println(filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	var starnewsArray []Starnew
	if err := json.Unmarshal(bytes, &starnewsArray); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	for _, starnews := range starnewsArray{
		fmt.Println(starnews.Create_time)
	}
	return starnewsArray, nil
}
