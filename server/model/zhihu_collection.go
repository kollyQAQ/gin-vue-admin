package model

type ZhihuCollection struct {
	ID       uint   `gorm:"primary_key"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Priority string `json:"priority"`
	Author   string `json:"author"`
}

func (ZhihuCollection) TableName() string {
	return "t_zhihu_collection"
}

type ZhihuCollectionWithContent struct {
	ID       uint   `gorm:"primary_key"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Priority string `json:"priority"`
	Author   string `json:"author"`
	Remark   string `json:"remark"`
}

func (ZhihuCollectionWithContent) TableName() string {
	return "t_zhihu_collection"
}
