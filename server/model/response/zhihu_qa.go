package response

import "gin-vue-admin/model"

type ZhihuQaResponse struct {
	Qa model.ZhihuQuestionAnswer `json:"qa"`
}

type ZhihuQaListResponse struct {
	GoodsList []model.ZhihuGoods `json:"goods_list"`
}
