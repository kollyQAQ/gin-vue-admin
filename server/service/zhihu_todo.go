package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func GetTodoList() (err error, list []*model.ZhihuTodo) {

	err = global.GVA_DB.Model(&model.ZhihuTodo{}).Find(&list).Error

	return err, list
}
