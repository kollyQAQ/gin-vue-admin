package response

import "gin-vue-admin/model"

type ZhihuStatResponse struct {
	Stat   ZhihuStat            `json:"stat"`
	QaStat []*model.ZhihuQaStat `json:"qa_stat"`
}

type ZhihuStat struct {
	TodayFee      int `json:"today_fee"`
	NoFeeGoodsNum int `json:"no_fee_goods_num"`
	TodoNumber    int `json:"todo_number"`
	TodayView     int `json:"today_view"`
}
