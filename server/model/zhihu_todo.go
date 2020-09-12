package model

type ZhihuTodo struct {
	ID     uint   `gorm:"primary_key"`
	Text   string `json:"text"`
	Status int    `json:"status"`
	Done   bool   `json:"done" gorm:"-"`
	Remark string `json:"remark"`
}

func (ZhihuTodo) TableName() string {
	return "t_todo"
}
