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
		//cmd := exec.Command("python36", "/root/pickme/backend/informarion_spider/starspider.py", "--neteasyid="+args.Account_id,"--name="+args.Star_name)
		cmd := exec.Command("ls")
		fmt.Println(cmd)
		var out bytes.Buffer

		cmd.Stdout = &out
		err = cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s", out.String())
		jsonStr := out.String()
		jsonStr = `[{'title': '《快乐大本营》王俊凯情商逆天, 一句话得罪全场, 一句话挽回场面', 'url': 'http://3g.163.com/idol/article/DQ8CUP3505370U9R.html', 'img': 'http://dingyue.nosdn.127.net/TNGvyEFN=hpZnbr1CWvfI0UJLEyoV35TzcJjEhSU23YCI1535375890631.jpg', 'create_time': '2018-08-27 21:18:42', 'source': '神马大娱乐'}, {'title': '舒淇让王俊凯多吃红烧肉，白举纲一句话让众人笑翻，网友：太逗了', 'url': 'http://3g.163.com/idol/article/DQ8CNQVD05370NUB.html', 'img': 'http://dingyue.nosdn.127.net/jI2e27qLFqkrNcKELXefZ3o4==amD7tq0rcbhQZbtwhr51535375667453.jpg', 'create_time': '2018-08-27 21:14:48', 'source': '娱乐心报 道'}, {'title': '《天坑鹰猎》最新海报太搞笑了，王俊凯火炬烧手？文淇左脚失踪？', 'url': 'http://3g.163.com/idol/article/DQ7T6LD3051788IF.html', 'img': 'http://dingyue.nosdn.127.net/7U8z5zCAPAslzE=KYYCIIBKY30Lo8to6FcQNlJWN08qbS1535359243106.jpg', 'create_time': '2018-08-27 20:25:01', 'source': '娱乐大爆炸'}, {'title': '时代变得真快啊! 王俊凯都开始拍床戏了!', 'url': 'http://3g.163.com/idol/article/DQ88CI1105370899.html', 'img': 'http://dingyue.nosdn.127.net/HTqOdwduMZMz2OBCeLKL6OELZUdKb8oqsQ5WzgiZOk7ve1535371120293.jpg', 'create_time': '2018-08-27 19:58:43', 'source': '娱乐小能手啊'}, {'title': '和明星做同学是什么体验？哎呦，心疼那个和王俊凯对视后的女生！', 'url': 'http://3g.163.com/idol/article/DQ86JMVU0537121N.html', 'img': 'http://dingyue.nosdn.127.net/mwf=82YJRS2K51FoCeuLwKvcZlpgaEXpFX4PMKyj3wkZG1535369238519.gif', 'create_time': '2018-08-27 19:27:46', 'source': '艳芳看娱乐'}, {'title': '《中餐厅》最暖心的一幕, 苏有朋和王俊凯的举动引网友怒赞', 'url': 'http://3g.163.com/idol/article/DQ86EHBP0521KOEV.html', 'img': 'http://dingyue.nosdn.127.net/zRp289KOJv97Y5F0nANaY=eex=8ogOXJUGmzzTsAV0IjE1535369090467compressflag.png', 'create_time': '2018-08-27 19:24:55', 'source': '文化强载'}, {'title': '王俊凯激动到把鞋都掉了，李维嘉赶忙去捡，有谁注意吴昕！', 'url': 'http://3g.163.com/idol/article/DQ84V8UM0517MA7I.html', 'img': 'http://dingyue.nosdn.127.net/Iny=SQ5vuE90GKyvaT87YhPoR88r9xw8PoetyYgzPGlED1535367498434compressflag.jpg', 'create_time': '2018-08-27 18:59:10', 'source': '小小泡腾 片'}, {'title': '快本：王俊凯最亲的人不是爸妈！舒淇无意间透露白举纲的真实家境', 'url': 'http://3g.163.com/idol/article/DQ83PRUI0517T7CE.html', 'img': 'http://dingyue.nosdn.127.net/U3d54KNGFUOqt7V9MR8yrSyXXlhkwDbEKUI0lREhUSXa21535366309511.jpeg', 'create_time': '2018-08-27 18:38:39', 'source': '天佑 娱乐圈'}, {'title': '王俊凯迎来荧幕初吻，画面却被吐槽太丑，网友：只是缺乏经验！', 'url': 'http://3g.163.com/idol/article/DQ83CR9T0518N5OU.html', 'img': 'http://dingyue.nosdn.127.net/7IMLSojzm96UMTzmHob3Vt6voIOCxvO58LYnmebEmeiu51535365888913compressflag.png', 'create_time': '2018-08-27 18:31:34', 'source': '型男时装控'}, {'title': '天坑鹰猎的预告片！男主王俊凯', 'url': 'http://3g.163.com/idol/article/DQ8331L205249LJB.html', 'img': 'http://dingyue.nosdn.127.net/lTOcPOUwhrmIZs1SMBJNXVOkAp=enAdeN4t8kdpjwJVPl1535365534718compressflag.jpg', 'create_time': '2018-08-27 18:26:12', 'source': '娱乐旅行事'}, {'title': '周年庆时：王俊凯主动跟千玺互动，结果千玺的反应让粉丝委屈极了', 'url': 'http://3g.163.com/idol/article/DQ82MKRM0517ARUK.html', 'img': 'http://dingyue.nosdn.127.net/LTyO01T4iSTNFyPLYclol3lWNk0rD3eBOB9U0nV2fDX3V1535365105720compressflag.jpg', 'create_time': '2018-08-27 18:19:25', 'source': 'TF娱乐'}, {'title': '《天坑鹰猎》王俊凯、文淇开篇高甜互动, 打架变“床咚”', 'url': 'http://3g.163.com/idol/article/DQ822K960517O17R.html', 'img': 'http://dingyue.nosdn.127.net/0mKDrGnJDVQS8x8atN6Pvg4geDuWHe3womw0EtjAKfl8g1535364488652.jpg', 'create_time': '2018-08-27 18:08:29', 'source': '明星娱乐哒帝'}, {'title': '《中餐厅》这些演员除了会演戏更会做饭，王俊凯竟然是厨房舞蹈家', 'url': 'http://3g.163.com/idol/article/DQ81L1NQ0517LJ5N.html', 'img': 'http://dingyue.nosdn.127.net/kji1KBZLcnpe1mEuW=vqVpK4pqw22G247bpt3=ZAB9I0Y1535363941214.jpg', 'create_time': '2018-08-27 18:01:09', 'source': '大咖頭条'}, {'title': '王俊凯新剧第一集就床咚女主，惨被骂流氓，粉丝要哭晕在厕所了', 'url': 'http://3g.163.com/idol/article/DQ80N4FQ05370PNR.html', 'img': 'http://dingyue.nosdn.127.net/V20GqGL=nKhaXPiZ2hQFOO3fLGY9pPDIwDOCmbydmL33i1535363066759compressflag.jpg', 'create_time': '2018-08-27 17:44:44', 'source': '快乐说娱乐'}, {'title': '粉丝大闹王俊凯吻戏床咚，导演解释借位拍摄，实则却为宣传手段？', 'url': 'http://3g.163.com/idol/article/DQ80BQSN05371VXP.html', 'img': 'http://dingyue.nosdn.127.net/g1mxOswRJgZIR6Y4KpxkQhN7aRJFm9P7IC69O8P8DvzQC1535362705624.jpg', 'create_time': '2018-08-27 17:38:32', 'source': '迪迪说娱'}, {'title': '国际秀场上的热巴、王俊凯、范爷、李宇春，谁更胜一筹？', 'url': 'http://3g.163.com/idol/article/DQ7VQRHK0517FA59.html', 'img': 'http://dingyue.nosdn.127.net/CVgIUNXp5cT7HbHPXwEwVTrS8RZHmKRclVUMCFTjVOHUw1535362131548.jpeg', 'create_time': '2018-08-27 17:29:16', 'source': '乃心'}, {'title': '她是和赵薇一样任性的“公主”，如今39岁成为王俊凯的时尚辣妈', 'url': 'http://3g.163.com/idol/article/DQ7VQAQK0517L7NO.html', 'img': 'http://dingyue.nosdn.127.net/0pe4pInnSuSbpQvOAbO=s2qufz9Jt1ncsG6hvWpXHW1PW1535362127908.jpeg', 'create_time': '2018-08-27 17:29:04', 'source': '李锐'}, {'title': '王俊凯十九字回怼赵薇，秒变专业“拆台凯”网友：大佬，不敢惹！', 'url': 'http://3g.163.com/idol/article/DQ7VJJME0529R7DQ.html', 'img': 'http://dingyue.nosdn.127.net/Qcp5OzZqf4oPOEKKkI0lTzjOQDtICvP0gaBsyZjDISEOB1535361912016.jpg', 'create_time': '2018-08-27 17:25:18', 'source': '疯狂搞笑的先生'}, {'title': '王俊凯太瘦了，外套宽的还能塞下一个人，往下看流苏想剪掉', 'url': 'http://3g.163.com/idol/article/DQ7S25U00517TGC9.html', 'img': 'http://dingyue.nosdn.127.net/A6bZ4QVlR4JOsp1TsIVlKiZh3ncaaboXSm9HPFzDuB8u=1535358186141compressflag.jpg', 'create_time': '2018-08-27 16:23:27', 'source': '娱乐首席评论员'}, {'title': '王俊凯很少吃红烧肉，苏有朋很爱吃，白举纲的话舒淇笑到飙泪', 'url': 'http://3g.163.com/idol/article/DQ7RPKDD05370WB7.html', 'img': 'http://dingyue.nosdn.127.net/Uptounsdn72vl4f7ufpBG6j3q5b3693tSqSoOH6oBz21Z1535357914543.jpg', 'create_time': '2018-08-27 16:18:42', 'source': '凝雁学堂'}, {'title': '王俊凯白举纲深情献歌，谁注意小凯最后的手势？网友：好想哭！', 'url': 'http://3g.163.com/idol/article/DQ7RHBU50525EQLL.html', 'img': 'http://dingyue.nosdn.127.net/pE1rmDF0d351Ug5FR2YDND1RlRc26iQbqd8gH5j3jky=N1535357622197.jpg', 'create_time': '2018-08-27 16:14:15', 'source': '端木侃生活'}, {'title': '王俊凯很暖，谢依霖抱大腿？他却热心于公益事业', 'url': 'http://3g.163.com/idol/article/DQ7QNGV505371X8M.html', 'img': 'http://dingyue.nosdn.127.net/zxzIw73YQA9vVUUXeQgH77tao6TIr0DGRDG4x6R1wmma31535356789812.jpg', 'create_time': '2018-08-27 16:00:05', 'source': '莫昔阳'}, {'title': '《中餐厅》厨房都是苍蝇，王俊凯苏有朋看见之后的反应让人无语', 'url': 'http://3g.163.com/idol/article/DQ7QDFLV053716JE.html', 'img': 'http://dingyue.nosdn.127.net/YgX45FaUAnusCWHD6gvxTLroGCaCuup5XYPMdHLeJRd941535356360684.jpg', 'create_time': '2018-08-27 15:54:36', 'source': '莜面娱乐圈'}, {'title': '王俊凯话筒声音小？依旧努力演唱，粉丝落泪', 'url': 'http://3g.163.com/idol/article/DQ7Q02SI0517O4KN.html', 'img': 'http://dingyue.nosdn.127.net/YVaxty6ovv3OtmOvVYqgFqzwldQ1Jhfxx8C=cbCyiku7S1535355992774.jpg', 'create_time': '2018-08-27 15:47:16', 'source': '大亨娱乐家'}, {'title': '少年初成长 ，王俊凯“天坑鹰猎”来袭！电影院走起！', 'url': 'http://3g.163.com/idol/article/DQ7PI1J2053726DZ.html', 'img': 'http://dingyue.nosdn.127.net/eHzv2qR5ZmOsjr3OKWr0qRgdVPXN8=IU3T7AJPPzakbU21535355472161.jpeg', 'create_time': '2018-08-27 15:39:34', 'source': '要上天的晓晓'}, {'title': '王俊凯吃鸡腿直接上 手抓，背后舒淇的眼神亮了，网友：偶像包袱呢', 'url': 'http://3g.163.com/idol/article/DQ7PF0N20525LCN7.html', 'img': 'http://dingyue.nosdn.127.net/7O2oBDNRQ5yTaQtJPMtWGCNkh5QDecFacNZAlouIygEN51535355465707.jpg', 'create_time': '2018-08-27 15:38:02', 'source': '勇山微元素'}, {'title': '那英为王俊凯男主大戏《天坑鹰猎》献唱主题曲', 'url': 'http://3g.163.com/idol/article/DQ7PTA0D0517DQF0.html', 'img': 'http://crawl.nosdn.127.net/1b945890185d6231498db95a9582b16e.jpg', 'create_time': '2018-08-27 15:30:49', 'source': '环球网娱乐'}, {'title': '高能！快本王俊凯向何老师撒娇，何炅笑到模糊：没法回头看你！', 'url': 'http://3g.163.com/idol/article/DQ7OC7280517TOV3.html', 'img': 'http://dingyue.nosdn.127.net/6dpssG4LwRHl5TAZLAolV25NoHdky1zH4wRKJXrzOzLGm1535354304850.jpg', 'create_time': '2018-08-27 15:18:57', 'source': '长腿腿哥'}, {'title': '王俊凯和舒淇对何老师撒娇，一个可爱，一个想把所有东西都给她', 'url': 'http://3g.163.com/idol/article/DQ7O873D0517WT0B.html', 'img': 'http://dingyue.nosdn.127.net/WLFSZ82RtdTkw0tmOLdq==kOflYzOoeIkHxKnbgFU8o5a1535354181973.jpg', 'create_time': '2018-08-27 15:16:50', 'source': '野兽'}, {'title': '王俊凯五周年未播个人吉他solo，看完视频的小螃蟹表示：生气又心疼！', 'url': 'http://3g.163.com/idol/article/DQ7NSSND0517R23S.html', 'img': 'http://dingyue.nosdn.127.net/CH0I1wj2acfEaBAeeDQB=fW08=8D8AJICtZeGnBTYPonp1535353831687.jpg', 'create_time': '2018-08-27 15:10:34', 'source': '读小宝娱乐'}, {'title': '外国客人挑战红辣椒，王俊凯皮一下不忘补刀', 'url': 'http://3g.163.com/idol/article/DQ7NHF1B0517AJMK.html', 'img': 'http://crawl.nosdn.127.net/1e588fc7cffbed98b1861129e3768918.jpg', 'create_time': '2018-08-27 14:59:28', 'source': '娱乐广播网'}, {'title': '明星瞎说大实话：秦岚却哭穷，李易峰炫富，王俊凯也太诚实了', 'url': 'http://3g.163.com/idol/article/DQ7N2L8H0511LCPP.html', 'img': 'http://dingyue.nosdn.127.net/QsJF6o5UJT=9MUyjkhaQWtFzdJh3BfwDs4CKycHx83I0b1535352956043.jpg', 'create_time': '2018-08-27 14:56:14', 'source': '天高与我飞'}, {'title': '霍建华、古天乐、王俊凯，谁才是“古装第一美男”', 'url': 'http://3g.163.com/idol/article/DQ7MOKQB0517X01P.html', 'img': 'http://dingyue.nosdn.127.net/qBOV6QBNARdjXY3l4x4=O93A4v2buW7kM=OmJZubjcJDB1535352633949.jpeg', 'create_time': '2018-08-27 14:50:46', 'source': '清夏蓝鲸'}, {'title': '本以为王俊凯在《天坑鹰猎 》中会是个王者，没想到却是青铜', 'url': 'http://3g.163.com/idol/article/DQ7LOVES0517DR21.html', 'img': 'http://dingyue.nosdn.127.net/xwDDqE7adiCewtMOIQXP3oP=gYUfupKuR5779em3BT1411535351547256compressflag.jpg', 'create_time': '2018-08-27 14:33:34', 'source': 'TF娱乐圈'}, {'title': '王俊凯新歌取名有深意, 与白举纲合作太惊艳! 粉丝很好奇写歌过程', 'url': 'http://3g.163.com/idol/article/DQ7LF8P80521KOB5.html', 'img': 'http://dingyue.nosdn.127.net/eje1N6LSBrvHm0Fkes3dgjUOQITEoe8=7V=5Veq0N=tpD1535351288761.jpg', 'create_time': '2018-08-27 14:28:13', 'source': '无厘头文化'}, {'title': '1岁出道，戏中演 过李小璐王丽坤童年，17岁搭档王俊凯让人羡慕', 'url': 'http://3g.163.com/idol/article/DQ7KU4EQ05371NLT.html', 'img': 'http://dingyue.nosdn.127.net/HQWJ6yRLRGhOSCqEQABduqofKQFviBCYnPpTPUWdzdUy11535350581366compressflag.jpg', 'create_time': '2018-08-27 14:18:55', 'source': 'Blaire左撇子'}, {'title': '《天坑鹰猎》预告出炉，王俊凯的一些特写镜头让粉丝疯狂截图', 'url': 'http://3g.163.com/idol/article/DQ7JUJ9S0517DR21.html', 'img': 'http://dingyue.nosdn.127.net/QtMp4FMU5IRpZJVhq=GKi2S9A=axOH71tYCnh2uSC3GEQ1535349595596compressflag.jpg', 'create_time': '2018-08-27 14:01:41', 'source': 'TF娱乐圈'}, {'title': '《快本》谢娜读家书，王俊凯完全不care，何炅忍不住diss：麻烦听一下', 'url': 'http://3g.163.com/idol/article/DQ78LDUT051788IF.html', 'img': 'http://dingyue.nosdn.127.net/jSE42OudZioZVwTCKMS0Yu69mJM=S9JCFhQhBMvmbx3CJ1535337836451compressflag.jpg', 'create_time': '2018-08-27 13:35:01', 'source': ' 娱乐大爆炸'}, {'title': '男艺人扮女装，王俊凯最萝莉，李易峰最接地气，只有他最辣眼', 'url': 'http://3g.163.com/idol/article/DQ7FBA2T0528VQR3.html', 'img': 'http://dingyue.nosdn.127.net/j9vD6n05Vatt6hrACS0A1NGunifFQEjEReiz7I0lb59VK1535344857685compressflag.jpeg', 'create_time': '2018-08-27 12:41:09', 'source': '情感伦理'}, {'title': '王俊凯李维嘉同样在吃东西，何炅为何只吼王俊凯，却不吼李维嘉？', 'url': 'http://3g.163.com/idol/article/DQ7BOKTP05370H2Q.html', 'img': 'http://dingyue.nosdn.127.net/5uWih5rYAHK3MU21CCir5qcPvFG4ryKT8Ibc8V22Iy=E11535341104112.jpg', 'create_time': '2018-08-27 11:38:32', 'source': '魔心电影欣赏'}, {'title': '《快乐大本营》戴这款表的姑娘 一定是王俊凯的铁粉', 'url': 'http://3g.163.com/idol/article/DQ75MN0S051892T4.html', 'img': 'http://dingyue.nosdn.127.net/QWmpbK0vhu7TcSWz71x7TvDzdRUVOYyRdwgcZU3KzfqAP1535334737107.jpg', 'create_time': '2018-08-27 09:52:37', 'source': '表迷'}, {'title': '王俊凯私下昵称苏有朋“爸爸”，苏有朋说：第一次看他就觉得像我', 'url': 'http://new.qq.com/omn/20180826/20180826A1G7ET.html', 'img': 'https://inews.gtimg.com/newsapp_ls/0/5005764376_640330/0', 'create_time': '2018-08-27 07:30:21', 'source': '腾讯娱乐'}, {'title': '王俊凯白举纲为一碗米粉的做法产生分歧，舒淇巧妙解围尽显情商', 'url': 'http://3g.163.com/idol/article/DQ63H1VP0517C26T.html', 'img': 'http://dingyue.nosdn.127.net/R2mQ18IQ5wpFJeZQKVbyk5qKaDneVnrmLcZwffz3dwkIv1535298034447compressflag.jpg', 'create_time': '2018-08-26 23:55:20', 'source': '美剧集中营'}, {'title': '电视剧《遮天》定妆完毕, 王俊凯、关晓彤任其主角, 胡歌客串出场!', 'url': 'http://3g.163.com/idol/article/DQ60VBHB05370G7W.html', 'img': 'http://dingyue.nosdn.127.net/AJIigGdWQ=KiWzIHbw4dxEVnCMN2UHjj0X0Gnd0dHbikg1535296225674.jpg', 'create_time': '2018-08-26 23:10:43', 'source': '广彬娱乐'}, {'title': '白敬 亭资源？王俊凯献“荧幕初吻”？冯绍峰赵丽颖在一起了吗？', 'url': 'http://3g.163.com/idol/article/DQ60DBF90517O3JI.html', 'img': 'http://dingyue.nosdn.127.net/HhB8HJn4UkDh3z34o4HAaFJ5QFtim4UmlpNe6srKuAJqd1535295602382.jpg', 'create_time': '2018-08-26 23:00:59', 'source': '娱人物'}, {'title': '白举纲说不辣，王俊凯真的很辣！《中餐厅》小白疯狂吸粉', 'url': 'http://3g.163.com/idol/article/DQ5VLVQT05370O3A.html', 'img': 'http://dingyue.nosdn.127.net/XEBRJtP4j1UnAwtXi3N2c4jt8XawWQgnQoYyTDnruIpcg1535294845778compressflag.jpg', 'create_time': '2018-08-26 22:48:08', 'source': '不二侃娱乐'}, {'title': '白举纲就在舒淇面前，她却只给王俊凯夹鸡腿，小白的样子令人心疼', 'url': 'http://3g.163.com/idol/article/DQ5VCFBS05371T7L.html', 'img': 'http://dingyue.nosdn.127.net/APPKHg0LcgA6oohnl3pVivXO6EMZ1tVfOGFJeBnRLgKxk1535294564774.jpg', 'create_time': '2018-08-26 22:42:59', 'source': '碧兰爱娱乐'}, {'title': '赵薇喂王俊凯吃美食, 李沁邓伦同用一双筷子, 明星私下感情真好', 'url': 'http://3g.163.com/idol/article/DQ5UB0G30514KNSK.html', 'img': 'http://dingyue.nosdn.127.net/7winroFZHCPtKdQdDTO4IayGuFlxvPaKN8j0U=3fNlk6U1535293465163compressflag.jpg', 'create_time': '2018-08-26 22:24:39', 'source': '听雨声营养师'}, {'title': '《中餐厅》赵薇做菜失误多出一份，谁留意王俊凯推销穿帮', 'url': 'http://3g.163.com/idol/article/DQ5SVLAT0517T9SE.html', 'img': 'http://dingyue.nosdn.127.net/WWNHZOGOB4=tPaOK7bCJsQqGkcmZvKShYqf0aUsJ67q2m1535292051972compressflag.jpg', 'create_time': '2018-08-26 22:19:09', 'source': '妖麟娱 乐秀'}, {'title': '《中餐厅2》小孩叫王俊凯叔叔，叫苏有朋爷爷，称呼赵薇更调皮', 'url': 'http://3g.163.com/idol/article/DQ5T9B0L0517T9SE.html', 'img': 'http://dingyue.nosdn.127.net/Omq=oiCuIJPlwmgvoleOjwxz6oo6AwYznS5PqcJO=0RSN1535292370343compressflag.jpg', 'create_time': '2018-08-26 22:17:16', 'source': '妖麟娱乐秀'}]`
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
				fmt.Println("insert",v)
				sqltool.StarsuckEngine.Insert(v)

			}else{
				fmt.Println("has",v)
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
