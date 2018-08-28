package news_states

import (
	"time"
)

type News struct{
	News_id			int			`xorm:"not null pk autoincr INT(11)" json:"news_id"`
	Star_id			int			`json:"star_id"`
	Img				string		`json:"img"`
	Title			string		`json:"title"`
	News_url		string		`json:"news_url"`
	Source			string		`json:"source"`
	Create_time		string		`json:"create_time"`
}


type Agenda struct {
	Agenda_id   int       `xorm:"not null pk autoincr INT(11)" json:"agenda_id"`
	Star_id     int       `json:"star_id"`
	Detail_time time.Time `json:"detail_time"`
	Location    string    `json:"location"`
	Content     string    `json:"content"`
}

type State struct {
	State_id		int			`xorm:"not null pk autoincr INT(11)" json:"state_id"`
	Account_name	string		`json:"account_name"`
	Content			string		`json:"content"`
	Create_time		time.Time	`json:"create_time"`
	Imgs			string		`json:"imgs"`
	Source			string		`json:"source"`
}

type info_source struct {
	Info_id 		int 		`xorm:"not null pk autoincr INT(11)" json:"info_id"`
	Star_id 		int 		`json:"star_id"`
	Source          string      `json:"source"`
	Account_name    string      `json:"account_name"`
	Usage_type      int         `json:"usage_type"` //1:state,2:news',
}