package model

type ZhihuQuestionAnswer struct {
	ID              int    `gorm:"primary_key"`
	Qid             string `json:"qid"`
	Type            int    `json:"type"`
	Qname           string `json:"qname"`
	ViewTotal       int    `json:"view_total"`
	AnswerTotal     int    `json:"answer_total"`
	TodayAddAnswer  int    `json:"today_add_answer"`
	TodayAddView    int    `json:"today_add_view"`
	ThreeDayAddView int    `json:"three_day_add_view"`
	Aid             string `json:"aid"`
	Rank            string `json:"rank"`
	LikeNum         string `json:"like_num"`
	WithCard        string `json:"with_card"`
	UpdateTime      string `json:"update_time"`
}

func (ZhihuQuestionAnswer) TableName() string {
	return "view_question_answer"
}

type ZhihuQuestion struct {
	ID          int    `gorm:"primary_key"`
	Type        int    `json:"type"`
	Qid         string `json:"qid"`
	Title       string `json:"title"`
	ViewTotal   int    `json:"view_total"`
	WatchTotal  int    `json:"watch_total"`
	AnswerTotal int    `json:"answer_total"`
}

func (ZhihuQuestion) TableName() string {
	return "t_zhihu_question"
}

type ZhihuAnswer struct {
	ID       int    `gorm:"primary_key"`
	Qid      string `json:"qid"`
	Aid      string `json:"aid"`
	Rank     string `json:"rank"`
	LikeNum  int    `json:"like_num"`
	WithCard string `json:"with_card"`
}

func (ZhihuAnswer) TableName() string {
	return "t_zhihu_answer"
}

type ZhihuQuestionHistory struct {
	Date    string `json:"date"`
	Qid     string `json:"qid"`
	Qname   string `json:"qname"`
	AddView int    `json:"add_view"`
}

func (ZhihuQuestionHistory) TableName() string {
	return "view_question_history"
}
