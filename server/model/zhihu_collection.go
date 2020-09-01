package model

type ZhihuCollection struct {
	ID       uint   `gorm:"primary_key"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Priority string `json:"priority"`
	Author   string `json:"author"`
	Remark   string `json:"remark"`
}

func (ZhihuCollection) TableName() string {
	return "t_zhihu_collection"
}
