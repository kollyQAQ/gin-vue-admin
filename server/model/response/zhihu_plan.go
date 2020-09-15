package response

import "gin-vue-admin/model"

type ZhihuPlanResponse struct {
	Todo  []*model.ZhihuPlan `json:"todo"`
	Inbox []*model.ZhihuPlan `json:"inbox"`
	Doing []*model.ZhihuPlan `json:"doing"`
	Done  []*model.ZhihuPlan `json:"done"`
}
