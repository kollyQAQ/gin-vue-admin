package request

import "gin-vue-admin/model"

// 分页条件查询及排序结构体
type SearchQaParams struct {
	model.ZhihuQuestionAnswer
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}

type CreateQaParams struct {
	Qid      string `json:"qid"`
	Qname    string `json:"qname"`
	Aid      string `json:"aid"`
	Type     int    `json:"type"`
	WithCard string `json:"with_card"`
}

type CreateQuestionParams struct {
	UserID uint   `json:"user_id"`
	Qid    string `json:"qid"`
	Type   int    `json:"type"`
}

type UpdateAnswerParams struct {
	UserID   uint   `json:"user_id"`
	Qid      string `json:"qid"`
	Aid      string `json:"aid"`
	WithCard string `json:"with_card"`
}

type GetQaByQidParams struct {
	Qid    string `json:"qid"`
	UserID int    `json:"user_id"`
}
