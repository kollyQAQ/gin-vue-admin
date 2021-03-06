package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"strconv"
)

func GetTodoList(userID uint) (err error, list []*model.ZhihuTodo) {

	err = global.GVA_DB.Model(&model.ZhihuTodo{}).Where("user_id = ?", userID).Find(&list).Error

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

func GetPlanList() (err error, list []*model.ZhihuPlan) {
	err = global.GVA_DB.Model(&model.ZhihuPlan{}).Find(&list).Error

	return err, list
}

func AddPlan(plan model.ZhihuPlan) (err error) {
	return global.GVA_DB.Create(&plan).Error
}

func UpdatePlan(status string, plan []*model.ZhihuPlan) (err error) {
	for _, plan := range plan {
		plan.Status, _ = strconv.Atoi(status)
		err = global.GVA_DB.Save(&plan).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTypeList(userID uint) (err error, list []*model.ZhihuType) {

	err = global.GVA_DB.Model(&model.ZhihuType{}).Where("user_id = ?", userID).Find(&list).Error

	return err, list
}
