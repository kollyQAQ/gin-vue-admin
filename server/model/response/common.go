package response

type PageResult struct {
	List     interface{} `json:"list"`
	Category []*Category `json:"category"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Category struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
