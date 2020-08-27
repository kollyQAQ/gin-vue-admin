package model

type ZhihuGoods struct {
	ID            uint    `gorm:"primary_key"`
	SkuID         string  `json:"sku_id"`
	SkuName       string  `json:"sku_name"`
	Price         float64 `json:"price"`
	FeeRate       float64 `json:"fee_rate"`
	Fee           float64 `json:"fee"`
	JdSale        uint    `json:"jd_sale"`
	ThirdCategory string  `json:"third_category" gorm:"column:cid3_name"`
	UpdateTime    string  `json:"update_time"`
}

func (ZhihuGoods) TableName() string {
	return "t_zhihu_goods"
}

type ZhihuGoodsWithContent struct {
	ID      uint   `gorm:"primary_key"`
	SkuID   string `json:"sku_id"`
	SkuName string `json:"sku_name"`
	Content string `json:"content"`
}

func (ZhihuGoodsWithContent) TableName() string {
	return "t_zhihu_goods"
}

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

type ZhihuCollection struct {
	ID     uint   `gorm:"primary_key"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	Author string `json:"author"`
	Remark string `json:"remark"`
}

func (ZhihuCollection) TableName() string {
	return "t_zhihu_collection"
}
