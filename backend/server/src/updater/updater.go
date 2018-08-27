package updater

import (
	"os/exec"
	"fmt"
	"bytes"
	"log"
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
			log.Fatal(err)
		}
		//fmt.Printf("%s", out.String())
		jsonStr := out.String()
		//jsonStr = `[{"title": "\u6768\u8d85\u8d8a\u7684\u7537\u7c89\u4e1d\uff0c\u5e94\u8be5\u5bf9\u5434\u4ea6\u51e1\u6709\u4e86\u5168\u65b0\u8ba4\u8bc6\u4e86", "url": "http://3g.163.com/idol/article/DQ6287OA0517O137.html", "img": "http://dingyue.nosdn.127.net/LOpSrtcNgfOCnDI3ZxoT=D4=C5ZJmz40mAdAo55jbiCgz1535297568434compressflag.jpg", "create_time": "2018-08-26 23:33:02", "source": "\u6f47\u6d12\u5c0f\u751f\u5a31\u4e50"}, {"title": "\u5a31\u4e500826 \u5434\u4ea6\u51e1MMVAs\u5f69\u6392\u540e\u8bb0\uff1a\u8d8a\u52aa\u529b\u8d8a\u5e78\u8fd0", "url": "http://3g.163.com/idol/article/DQ60521505371GQR.html", "img": "http://dingyue.nosdn.127.net/usdDi9yCR6DVf7d173BXyf2rJjSXg40gYyb5g3DoiHZb=1535295369643compressflag.jpg", "create_time": "2018-08-26 22:56:26", "source": "\u8fea\u4e3d\u70ed\u62cd"}]`

		fmt.Println(jsonStr)
		news := make([]news_states.News,0)
		if err := json.Unmarshal([]byte(jsonStr), &news); err == nil {
			fmt.Println(len(news),news)
		} else {
			fmt.Println(err)
			continue
		}

		for _,v:=range(news){
			v.Star_id = args.Star_id
			yes, err := sqltool.StarsuckEngine.Table("news").Get(v)
			if err != nil {
				fmt.Println(err)
				break
			}
			if!yes {
				sqltool.StarsuckEngine.Insert(v)

			}else{
				break
			}
		}
	}

}

func StatesUpdater(){
	cmd := exec.Command("python", "/root/pickme/backend/informarion_spider/weiboSpider.py","--id=XXX","--filename='/root/pickme/backend/informarion_spider/fuck.json'")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())

	jsonStr := `[{"account_id": "6349794947", "content": "#\u738b\u4fca\u51ef\u98d8\u67d4\u5168\u7403\u4ee3\u8a00##\u98d8\u67d4\u5c11\u5e74\u738b\u4fca\u51ef#\u5c11\u5e74\u5929\u751f\u7075\u52a8\uff0c\u7738\u82e5\u707f\u661f\u6e29\u5982\u7389\u3002\u6b22\u8fce@TFBOYS-\u738b\u4fca\u51ef \u6b63\u5f0f\u6210\u4e3a\u98d8\u67d4\u5168\u7403\u4ee3\u8a00\u4eba\u3002\u4e16\u754c\u518d\u5927\u4e0d\u6539\u6e05\u6f88\u672c\u771f\uff0c\u4e3a\u98d8\u67d4\u6ce8\u5165\u65b0\u9c9c\u6d3b\u529b\uff0c\u4eca\u540e\u643a\u624b\u540c\u884c\u3002 \u200b\u200b", "source": "\u5fae\u535a", "create_time": "2018-08-27 09:21", "imgs": "http://wx3.sinaimg.cn/wap180/7218668fgy1funzfrxs5dj22o03quqvc.jpg", "account_name": "\u738b\u4fca\u51efKarryWang\u5de5\u4f5c\u5ba4"}, {"account_id": "6349794947", "content": "#\u738b\u4fca\u51ef[\u8d85\u8bdd]#@TFBOYS-\u738b\u4fca\u51ef \u6e5b\u6e05\u5982\u6c34\uff0c\u6717\u6717\u5fc3\u6674\u3002\u6c89\u674e\u6d6e\u74dc\uff0c\u590f\u65e5\u7ef5\u7ef5\u3002Karry\u2019s Album\uff0ccoming soon\u2026\u2026[\u5fae\u98ce][\u5fae\u98ce][\u5fae\u98ce] #karry\u738b\u7684\u65e5\u5e38# \u738b\u4fca\u51efKarryWang\u5de5\u4f5c\u5ba4\u7684\u79d2\u62cd\u89c6\u9891 \u200b\u200b", "source": "\u5fae\u535a", "create_time": "2018-08-26 13:21", "imgs": "\u65e0", "account_name": "\u738b\u4fca\u51efKarryWang\u5de5\u4f5c\u5ba4"}]`

	states := make([]news_states.State,0)
	if err := json.Unmarshal([]byte(jsonStr), &states); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(states)
	}

}
