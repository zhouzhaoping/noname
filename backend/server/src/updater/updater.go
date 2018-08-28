package updater

import (
	"os/exec"
	"fmt"
	"bytes"
	"encoding/json"
	"news_states"
	"sqltool"
)

type star_id_neteasyid struct {
	Star_id int		`json:"star_id"`
	Account_id string	`json:"account_id"`
	Star_name string 	`json:"star_name"`
}

func NewsUpdater() {

	u_n_list := make([]star_id_neteasyid, 0)

	err := sqltool.StarsuckEngine.Table("info_source").Cols("info_source.star_id","account_id","star_name").
		Join("INNER", "star_info","star_info.star_id = info_source.star_id").
		Where("source=?","netease").Find(&u_n_list)
	if err != nil {
		fmt.Println(err, "数据库错误")
		return
	} else {
		fmt.Println(len(u_n_list),u_n_list)
	}

	for _,args:=range(u_n_list){
		// python36 starspider.py --neteasyid=31 --name="王俊凯"
		cmd := exec.Command("python36", "/root/pickme/backend/informarion_spider/starspider.py", "--neteasyid="+args.Account_id,"--name="+args.Star_name)
		//cmd := exec.Command("ls")
		fmt.Println(cmd)
		var out bytes.Buffer

		cmd.Stdout = &out
		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
			continue
			//log.Fatal(err)
		}
		//fmt.Printf("%s", out.String())
		jsonStr := out.String()
		//jsonStr = `[{"title": "《快乐大本营》王俊凯情商逆天, 一句话得罪全场, 一句话挽回场面", "news_url": "http://3g.163.com/idol/article/DQ8CUP3505370U9R.html", "img": "http://dingyue.nosdn.127.net/TNGvyEFN=hpZnbr1CWvfI0UJLEyoV35TzcJjEhSU23YCI1535375890631.jpg", "create_time": "2018-08-27 21:18:42", "source": "神马大娱乐"}]`
		fmt.Println(jsonStr)
		news := make([]news_states.News,0)
		if err := json.Unmarshal([]byte(jsonStr), &news); err == nil {
			fmt.Println(len(news))//,news)
		} else {
			fmt.Println(err)
			continue
		}

		for _,v:=range(news){
			v.Star_id = args.Star_id
			yes, err := sqltool.StarsuckEngine.Table("news").Get(&v)
			if err != nil {
				fmt.Println(err)
				break
			}
			if!yes {
				fmt.Println("insert",v)
				sqltool.StarsuckEngine.Insert(v)

			}else{
				fmt.Println("has",v)
				break
			}
		}
	}

}

type star_id_weiboid struct {
	Star_id int		`json:"star_id"`
	Account_id string	`json:"account_id"`
	Source string  	`json:"source"`
}

func StatesUpdater(){

	s_w_list := make([]star_id_weiboid, 0)

	err := sqltool.StarsuckEngine.Table("info_source").Cols("info_source.star_id","account_id","source").
		Join("INNER", "star_info","star_info.star_id = info_source.star_id").
		Where("source=?","weibo").Find(&s_w_list)
	if err != nil {
		fmt.Println(err, "数据库错误")
		return
	} else {
		fmt.Println(len(s_w_list),s_w_list)
	}

	for _,args:=range(s_w_list){
		cmd := exec.Command("python", "/root/pickme/backend/informarion_spider/weiboSpider.py", "--user_id \""+args.Account_id+"\"", "--source \""+args.Source+"\"")
		//cmd := exec.Command("python", "/root/pickme/backend/informarion_spider/weiboSpider.py","--id=XXX","--filename='/root/pickme/backend/informarion_spider/fuck.json'")
		//cmd := exec.Command("ls")
		fmt.Println(cmd)
		var out bytes.Buffer

		cmd.Stdout = &out
		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
			continue
			//log.Fatal(err)
		}
		//fmt.Printf("%s", out.String())
		jsonStr := out.String()
		//jsonStr = `\[{"title": "《快乐大本营》王俊凯情商逆天, 一句话得罪全场, 一句话挽回场面", "url": "http://3g.163.com/idol/article/DQ8CUP3505370U9R.html", "img": "http://dingyue.nosdn.127.net/TNGvyEFN=hpZnbr1CWvfI0UJLEyoV35TzcJjEhSU23YCI1535375890631.jpg", "create_time": "2018-08-27 21:18:42", "source": "神马大娱乐"}]`
		fmt.Println(jsonStr)
		states := make([]news_states.News,0)
		if err := json.Unmarshal([]byte(jsonStr), &states); err == nil {
			fmt.Println(len(states),states)
		} else {
			fmt.Println(err)
			continue
		}

		for _,v:=range(states){
			v.Star_id = args.Star_id
			yes, err := sqltool.StarsuckEngine.Table("state").Get(&v)
			if err != nil {
				fmt.Println(err)
				break
			}
			if!yes {
				fmt.Println("insert",v)
				sqltool.StarsuckEngine.Insert(v)

			}else{
				fmt.Println("has",v)
				break
			}
		}
	}


}
