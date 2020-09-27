package request

import "gin-vue-admin/model"

// 分页条件查询及排序结构体
type UpdatePlanParams struct {
	Status string             `json:"status"`
	Data   []*model.ZhihuPlan `json:"data"`
}
