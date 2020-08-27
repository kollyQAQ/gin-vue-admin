package request

import "gin-vue-admin/model"

// 分页条件查询及排序结构体
type SearchCollectionParams struct {
	model.ZhihuCollection
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
