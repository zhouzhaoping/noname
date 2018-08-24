package star

type Star_info struct {
	Star_id 					int 	`xorm:"not null pk autoincr INT(11)" json:"star_id"`
	Star_name					string	`json:"star_name"`
	Img							string	`json:"img"`
	Average_rank				uint	`json:"average_rank"`
	Average_highest_rank		uint	`json:"average_highest_rank"`
	Baidu_index					uint	`json:"baidu_index"`
	Current_weibofans_num		uint	`json:"current_weibofans_num"`
	Yesterday_weibofans_num		uint	`json:"yesterday_weibofans_num"`
	Current_insfans_num			uint	`json:"current_insfans_num"`
	Yesterday_insfans_num		uint	`json:"yesterday_insfans_num"`
	Tvshow_num					uint	`json:"tvshow_num"`
	Ads_num						uint	`json:"ads_num"`
	Banner						string	`json:"banner"`
	Identify					string	`json:"identify"`
	Mv_num						int		`json:"mv_num"`
}


type Star_info_simple struct{
	Star_id		int		`json:"star_id"`
	Star_name	string	`json:"star_name"`
	Img 		string	`json:"img"`
}