package response

import "gin-vue-admin/model"

type ZhihuStatResponse struct {
	Stat       ZhihuStat            `json:"stat"`
	QaStat     []*model.ZhihuQaStat `json:"qa_stat"`
	SearchStat []*ZhihuSearchStat   `json:"search_stat"`
}

type ZhihuStat struct {
	TodayFee      int `json:"today_fee"`
	NoFeeGoodsNum int `json:"no_fee_goods_num"`
	TodoNumber    int `json:"todo_number"`
	TodayView     int `json:"today_view"`
}

type ZhihuSearchStat struct {
	ArticleID string `json:"article_id"`
	Remark    string `json:"remark"`
	Rank      int    `json:"rank"`
}
