package spider

type SkuResp struct {
	Response *GeneralRet `json:"jd_union_open_goods_promotiongoodsinfo_query_response"`
}

type GeneralRet struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type Result struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []*SkuData `json:"data"`
}

type SkuData struct {
	SkuID            int     `json:"skuId"`
	GoodsName        string  `json:"goodsName"`
	CommisionRatioPc float64 `json:"commisionRatioPc"` // 佣金比例
	InOrderCount     int     `json:"inOrderCount"`     // 订单数量
	IsJdSale         int     `json:"isJdSale"`         // 是否京东自营
	UnitPrice        float64 `json:"unitPrice"`        // 售价
}
