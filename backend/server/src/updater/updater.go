package updater

import (
	"os/exec"
	"fmt"
	"bytes"
	"encoding/json"
	"news_states"
	"sqltool"
	"time"
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

		time.Sleep(time.Minute * 5)

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
	Account_name string	`json:"account_name"`
	Account_id string	`json:"account_id"`
	Source string  	`json:"source"`
}

func StatesUpdater(){

	s_w_list := make([]star_id_weiboid, 0)

	err := sqltool.StarsuckEngine.Table("info_source").Cols("info_source.star_id","account_name","account_id","source").
		Join("INNER", "star_info","star_info.star_id = info_source.star_id").
		Where("source=? or source=?","weibo","instagram").Find(&s_w_list)
	if err != nil {
		fmt.Println(err, "数据库错误")
		return
	} else {
		fmt.Println(len(s_w_list),s_w_list)
	}

	for _,args:=range(s_w_list){
		cmd := exec.Command("python", "/root/pickme/backend/informarion_spider/weiboSpider.py", "--user_id="+args.Account_id, "--source="+args.Source)
		//cmd := exec.Command("python", "/root/test2.py", "--user_id","\""+args.Account_id+"\"", "--source","\""+args.Source+"\"")
		//cmd := exec.Command("python", "/root/pickme/backend/informarion_spider/weiboSpider.py","--id=XXX","--filename='/root/pickme/backend/informarion_spider/fuck.json'")
		//cmd := exec.Command("ls")
		fmt.Println(cmd)
		var out bytes.Buffer

		cmd.Stdout = &out
		err = cmd.Run()
		time.Sleep(time.Minute * 5)

		if err != nil {
			fmt.Println("wtf",err)
			fmt.Println(out.String())
			continue
			//log.Fatal(err)
		}

		fmt.Printf("%s", out.String())
		jsonStr := out.String()
		//jsonStr = `[{"account_id": "6349839494", "content": "#\u6211\u4e3a\u4e2d\u56fd\u822a\u5929\u6dfb\u71c3\u6599#\u6562\u4e8e\u505a\u68a6\uff0c\u52c7\u4e8e\u8ffd\u68a6\uff0c\u6bcf\u4e00\u4e2a\u5929\u9a6c\u884c\u7a7a\u90fd\u6709\u53ef\u80fd\u6210\u4e3a\u4e0b\u4e00\u4e2a\u5343\u8f7d\u96be\u9022\uff0c\u6211\u662f\u6613\u70ca\u5343\u73ba\uff0c\u8fd9\u662f\u6211\u7684\u592a\u7a7a\u68a6\u60f3\uff0c\u4f60\u7684\u5462\uff1fTFBOYS-\u6613\u70ca\u5343\u73ba\u7684\u79d2\u62cd\u89c6\u9891 \u200b\u200b", "source": "\u5fae\u535a", "create_time": "2018-08-27 11:30", "imgs": "\u65e0", "account_name": "\u6613\u70ca\u5343\u73baJacksonYee\u5de5\u4f5c\u5ba4"}, {"account_id": "6349839494", "content": "#\u98ce\u534e\u6063\u610f\u4e94\u5e74\u70bd\uff0c\u6613\u70ca\u5343\u73ba\u7834\u7acb\u65f6#\u4e0d\u66fe\u968f\u6ce2\u9010\u6d41\uff0c\u4ece\u672a\u653e\u6162\u811a\u6b65\u3002\u5c11\u5e74@TFBOYS-\u6613\u70ca\u5343\u73ba \u4e07\u822c\u6837\u8c8c\uff0c\u4e00\u822c\u6e29\u67d4\u3002#TFBOYS\u4e94\u5468\u5e74\u5f00\u59cb\u60f3\u8c61# \u8ffd\u5149\u706f\u4e0b\u7684\u821e\u53f0\u738b\u8005\uff0c\u76ee\u5149\u6240\u81f4\u7686\u662f\u4f60\u3002\u5e74\u5e74\u5c81\u5c81\u59cb\u7ec8\u966a\u4f34\uff0c\u5c81\u5c81\u5e74\u5e74\u6e29\u6696\u540c\u884c\u3002 \u200b\u200b", "source": "\u5fae\u535a", "create_time": "2018-08-25 15:30", "imgs": "http://wx1.sinaimg.cn/wap180/006VJhk2ly1fulz6gi8hej32kw3vc7wk.jpg", "account_name": "\u6613\u70ca\u5343\u73baJacksonYee\u5de5\u4f5c\u5ba4"}, {"account_id": "6349839494", "content": "\u90fd\u5230\u5bb6\u4e86\u4e48\uff1f\u4eca\u665a\u542c\u89c1\u4f60\u4eec\u7684\u58f0\u97f3\u4e86\uff0c\u8f9b\u82e6\u5927\u5bb6\uff0c\u65e9\u70b9\u4f11\u606f\u3002 \u200b\u200b", "source": "\u5fae\u535a", "create_time": "2018-08-25 01:43", "imgs": "http://wx2.sinaimg.cn/wap180/d7f7faddly1fulb9opq4vj23vc2kwnpe.jpg", "account_name": "\u6613\u70ca\u5343\u73baJacksonYee\u5de5\u4f5c\u5ba4"}]`
		fmt.Println(jsonStr)
		states := make([]news_states.State,0)
		if err := json.Unmarshal([]byte(jsonStr), &states); err == nil {
			fmt.Println(len(states),states)
		} else {
			fmt.Println(err)
			continue
		}

		for _,v:=range(states){
			v.Account_name = args.Account_name
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
