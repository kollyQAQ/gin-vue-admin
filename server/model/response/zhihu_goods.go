package response

import "gin-vue-admin/model"

type ZhihuGoodsResponse struct {
	Goods model.ZhihuGoodsWithContent `json:"goods"`
}

type ZhihuGoodsListResponse struct {
	GoodsList []model.ZhihuGoods `json:"goods_list"`
}
