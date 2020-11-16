package request

import "gin-vue-admin/model"

// 分页条件查询及排序结构体
type SearchGoodsParams struct {
	model.ZhihuGoods
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}

type SearchHighRateGoodsParams struct {
	model.ZhihuHighRateGoods
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
