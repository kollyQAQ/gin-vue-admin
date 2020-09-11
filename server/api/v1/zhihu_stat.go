package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

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
