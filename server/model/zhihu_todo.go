package model

type ZhihuTodo struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `json:"-"`
	Text   string `json:"text"`
	Status int    `json:"status"`
	Done   bool   `json:"done" gorm:"-"`
	Remark string `json:"remark"`
}

func (ZhihuTodo) TableName() string {
	return "t_todo"
}

type ZhihuPlan struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Type   int    `json:"type"`
	Text   string `json:"name"`
	Status int    `json:"status"`
	Level  int    `json:"level"`
	Remark string `json:"remark"`
}

func (ZhihuPlan) TableName() string {
	return "t_plan"
}

type ZhihuType struct {
	ID     uint   `gorm:"primary_key" json:"value"`
	UserID uint   `json:"-"`
	Label  string `json:"label"`
}

func (ZhihuType) TableName() string {
	return "t_zhihu_type"
}
