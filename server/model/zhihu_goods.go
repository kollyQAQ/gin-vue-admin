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
