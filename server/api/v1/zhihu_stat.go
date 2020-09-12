package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

func GetMusicList(c *gin.Context) {
	err, list := service.GetMusicList()

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(list, c)
	}
}

func GetTodoList(c *gin.Context) {
	err, list := service.GetTodoList()

	for _, item := range list {
		if item.Status == 1 {
			item.Done = true
		}
	}

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(list, c)
	}
}

func AddTodo(c *gin.Context) {
	var todo model.ZhihuTodo
	_ = c.ShouldBindJSON(&todo)

	err := service.AddTodo(todo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo model.ZhihuTodo
	_ = c.ShouldBindJSON(&todo)

	err := service.UpdateTodo(todo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo model.ZhihuTodo
	_ = c.ShouldBindJSON(&todo)

	err := service.DeleteTodo(todo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
