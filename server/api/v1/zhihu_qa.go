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

func GetQuestionAnswer(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchQaParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetZhihuQuestionAnswer(sp.ZhihuQuestionAnswer, sp.PageInfo, sp.OrderKey, sp.Desc)
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

func GetQaById(c *gin.Context) {
	var idInfo request.GetByIdStr
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, qa := service.GetQaById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.ZhihuQaResponse{Qa: qa}, c)
	}
}

func UpdateQa(c *gin.Context) {
	var qa model.ZhihuQuestionAnswer
	_ = c.ShouldBindJSON(&qa)
	ApiVerify := utils.Rules{
		"SkuID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(qa, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.UpdateAnswer(qa.Qid, qa.Aid, qa.WithCard)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

func CreateQa(c *gin.Context) {
	var question model.ZhihuQuestion
	_ = c.ShouldBindJSON(&question)
	ApiVerify := utils.Rules{
		"Qid": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(question, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.CreateQuestion(question)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteQa(c *gin.Context) {
	var a model.ZhihuQuestion
	_ = c.ShouldBindJSON(&a)
	//ApiVerify := utils.Rules{
	//	"ID": {utils.NotEmpty()},
	//}
	//ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	//if ApiVerifyErr != nil {
	//	response.FailWithMessage(ApiVerifyErr.Error(), c)
	//	return
	//}
	err := service.DeleteQuestion(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func QueryHistory(c *gin.Context) {
	var idInfo request.GetByIdStr
	_ = c.ShouldBindJSON(&idInfo)

	var res [][]interface{}
	dateList := []interface{}{"date"}
	viewList := []interface{}{"view"}
	//res = append(res, []interface{}{"date", "09-01", "09-02", "09-03", "09-04", "09-05", "09-06", "09-07"})
	//res = append(res, []interface{}{"view", 100, 1100, 1200, 2500, 2000, 400, 200})
	//response.OkWithData(res, c)

	err, list := service.QueryQuestionHistory(idInfo.Id)
	for _, history := range list {
		dateList = append(dateList, history.Date)
		viewList = append(viewList, history.AddView)
	}

	res = append(res, dateList)
	res = append(res, viewList)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(res, c)
	}
}

func QueryQaStat(c *gin.Context) {
	err, list := service.QueryQaStat()
	typeMap := service.GetQuestionTypeMap()
	for _, item := range list {
		item.TypeDesc = typeMap[item.Type]
	}

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(list, c)
	}
}
