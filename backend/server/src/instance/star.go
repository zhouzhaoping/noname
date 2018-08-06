package instance

type star_info struct {
	star_id 					int		`用户id`
	star_name					string	`姓名`
	img 						string	`头像`
	weibost_daily_rank			int		`微博超话日榜`
	asians_daily_rank			int		`亚洲新歌日榜`
	weibostar_power_daily_rank	int		`微博明星势力日榜`
	weixinstar_right_daily_rank	int		`微信明星权力日榜`
	weixinstar_power_daily_rank int		`微信明星势力日榜`
	average_rank 				int		`今日平均排名`
	average_highest_rank 		int		`平均历史最高名次`
	baidu_index 				int		`百度指数`
	current_weibofans_num 		int		`当日微博粉丝数`
	yesterday_weibofans_num 	int		`昨天粉丝数`
	current_insfans_num 		int		`当日ins粉丝数`
	yesterday_insfans_num 		int		`ins粉丝数`
	tvshow_num 					int		`已参加综艺节目数`
	ads_num 					int		`在接广告数`
}

type offical_account struct {
	account_id		int		`账户id `
	account_name	string	`账户名称 `
	star_id			int		`明星id`
	account_url		string	`链接`
}

