package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func GetMusicList() (err error, list []*model.ZhihuMusic) {

	err = global.GVA_DB.Model(&model.ZhihuMusic{}).Find(&list).Error

	return err, list
}
