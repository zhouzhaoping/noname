package handler

import (
	"time"
	"os/exec"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sort"
)

//结构体里的字段首字母必须大写，否则无法正常解析
type Starnew struct {
	Title string  			`json:"title"`
	Url string				`json:"url"`
	Img string				`json:"img"`
	Create_time string		`json:"create_time"`
	Source string			`json:"source"`
}
type starnewlist []Starnew

func (I starnewlist) Len() int {
	return len(I)
}
func (I starnewlist) Less(i, j int) bool {
	return I[i].Create_time > I[j].Create_time
}
func (I starnewlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
func MergeSlice(s1 []Starnew, s2 []Starnew) []Starnew {
	slice := make([]Starnew, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

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
func UpdateNews(starid string)[]byte{
	var starnewsArray []Starnew
	if starid == "follow" {
		starnewsArray0, _ := readFile("/root/pickme/backend/informarion_spider/news_list0.json")
		starnewsArray1, _ := readFile("/root/pickme/backend/informarion_spider/news_list1.json")
		starnewsArray2, _ := readFile("/root/pickme/backend/informarion_spider/news_list2.json")

		starnewsArray = MergeSlice(starnewsArray0[:2], starnewsArray1[0:2])
		starnewsArray = MergeSlice(starnewsArray, starnewsArray2[0:2])
		fmt.Println(len(starnewsArray))
		sort.Sort(starnewlist(starnewsArray))
	} else {//TODO 失败
		starnewsArray, _ := readFile("/root/pickme/backend/informarion_spider/news_list" + starid + ".json")
		starnewsArray = starnewsArray[:2]
		fmt.Println(len(starnewsArray))
		/*for _, starnews := range starnewsArray{
			fmt.Println(starnews)
		}*/
	}


	fmt.Println(starnewsArray)
	bytes, err := json.Marshal(starnewsArray)
	fmt.Println(string(bytes))
	if err != nil{
		fmt.Println(err.Error())
		return nil
	}else{
		fmt.Println(string(bytes))
	}
	return bytes
	//sqltool.StarsuckEngine
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

	//for _, starnews := range starnewsArray{
	//	fmt.Println(starnews.Create_time)
	//}
	//bytes, _ = json.Marshal(starnewsArray)
	//fmt.Println(string(bytes))
	fmt.Println(starnewsArray)
	return starnewsArray, err
}
