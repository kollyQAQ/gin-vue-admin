package model

type ZhihuQuestionAnswer struct {
	ID             uint   `gorm:"primary_key"`
	Type           int    `json:"type"`
	Qid            string `json:"qid"`
	Qname          string `json:"qname"`
	ViewTotal      uint   `json:"view_total"`
	AnswerTotal    uint   `json:"answer_total"`
	TodayAddAnswer int    `json:"today_add_answer"`
	TodayAddView   int    `json:"today_add_view"`
	Aid            string `json:"aid"`
	Rank           string `json:"rank"`
	LikeNum        string `json:"like_num"`
	WithCard       string `json:"with_card"`
	UpdateTime     string `json:"update_time"`
}

func (ZhihuQuestionAnswer) TableName() string {
	return "view_question_answer"
}
