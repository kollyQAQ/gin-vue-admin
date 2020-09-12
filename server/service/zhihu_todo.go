package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func GetTodoList() (err error, list []*model.ZhihuTodo) {

	err = global.GVA_DB.Model(&model.ZhihuTodo{}).Find(&list).Error

	return err, list
}

func UpdateTodo(todo model.ZhihuTodo) (err error) {
	err = global.GVA_DB.Save(&todo).Error

	return err
}

func AddTodo(todo model.ZhihuTodo) (err error) {
	return global.GVA_DB.Create(&todo).Error
}

func DeleteTodo(todo model.ZhihuTodo) (err error) {
	err = global.GVA_DB.Delete(todo).Error
	return err
}
