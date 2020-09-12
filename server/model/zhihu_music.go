package model

type ZhihuMusic struct {
	ID     uint   `gorm:"primary_key"`
	Mid    string `json:"mid"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
	Cover  string `json:"cover"`
	Lrc    string `json:"lrc"`
}

func (ZhihuMusic) TableName() string {
	return "t_music"
}
