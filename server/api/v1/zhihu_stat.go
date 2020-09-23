package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
)

func QueryLog(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchLogParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

	err, list, total := service.GetLogList(sp.ZhihuLog, sp.PageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		}, c)
	}
}

func QueryStat(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err1, list := service.QueryQaStat(waitUse.ID)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err1), c)
	}

	err2, data := service.QueryStat(waitUse.ID)
	if err2 != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err2), c)
	}

	response.OkWithData(resp.ZhihuStatResponse{
		QaStat: list,
		Stat:   data,
	}, c)

}

func GetMusicList(c *gin.Context) {
	err, list := service.GetMusicList()

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(list, c)
	}
}

func GetPlanList(c *gin.Context) {
	err, list := service.GetPlanList()

	var todo, doing, done, inbox []*model.ZhihuPlan
	for _, item := range list {
		if item.Status == 0 {
			inbox = append(inbox, item)
		} else if item.Status == 1 {
			todo = append(todo, item)
		} else if item.Status == 2 {
			doing = append(doing, item)
		} else if item.Status == 3 {
			done = append(done, item)
		}
	}

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(&resp.ZhihuPlanResponse{
			Inbox: inbox,
			Todo:  todo,
			Doing: doing,
			Done:  done,
		}, c)
	}
}

func GetTodoList(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, list := service.GetTodoList(waitUse.ID)

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

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	todo.UserID = waitUse.ID

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

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	todo.UserID = waitUse.ID

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

func GetTypeList(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	result := make([]*model.ZhihuType, 0)
	result = append(result, &model.ZhihuType{
		ID:     0,
		UserID: waitUse.ID,
		Label:  "未分类",
	})

	err, list := service.GetTypeList(waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		list = append(result, list...)
		response.OkWithData(list, c)
	}
}
