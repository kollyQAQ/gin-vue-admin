package model

type ZhihuGoods struct {
	ID            uint    `gorm:"primary_key"`
	UserID        uint    `json:"-"`
	SkuID         string  `json:"sku_id"`
	SkuName       string  `json:"sku_name"`
	Price         float64 `json:"price"`
	FeeRate       float64 `json:"fee_rate"`
	Fee           float64 `json:"fee"`
	JdSale        uint    `json:"jd_sale"`
	ThirdCategory string  `json:"third_category" gorm:"column:cid3_name"`
	OrderNum      uint    `json:"order_num"`
	UpdateTime    string  `json:"update_time"`
}

func (ZhihuGoods) TableName() string {
	return "t_zhihu_goods"
}

type ZhihuHighRateGoods struct {
	ID            uint    `gorm:"primary_key"`
	GoodsName     string  `json:"goods_name"`
	ShopName      string  `json:"shop_name"`
	Url           string  `json:"url"`
	CategoryOne   string  `json:"category_one"`
	CategoryTwo   string  `json:"category_two"`
	CategoryThree string  `json:"category_three"`
	Price         float64 `json:"price"`
	Fee           float64 `json:"fee"`
	FeeRate       float64 `json:"fee_rate"`
}

func (ZhihuHighRateGoods) TableName() string {
	return "t_jd_goods"
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
