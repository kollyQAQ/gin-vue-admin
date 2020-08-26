package model

type ZhihuGoods struct {
	ID            uint    `gorm:"primary_key"`
	SkuID         string  `json:"sku_id" gorm:"comment:'api路径'"`
	SkuName       string  `json:"sku_name" gorm:"comment:'api中文描述'"`
	Price         float64 `json:"price" gorm:"comment:'api组'"`
	FeeRate       float64 `json:"fee_rate" gorm:"default:'POST'" gorm:"comment:'方法'"`
	Fee           float64 `json:"fee" gorm:"default:'POST'" gorm:"comment:'方法'"`
	JdSale        uint    `json:"jd_sale"`
	ThirdCategory string  `json:"third_category" gorm:"column:cid3_name"`
}

// 设置User的表名为`profiles`
func (ZhihuGoods) TableName() string {
	return "t_zhihu_goods"
}
