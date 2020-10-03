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

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, list, total := service.GetZhihuQuestionAnswer(sp.ZhihuQuestionAnswer, sp.PageInfo, sp.OrderKey, sp.Desc, waitUse.ID)
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

func GetQaByID(c *gin.Context) {
	var idInfo request.GetByIdStr
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, qa := service.GetQaById(idInfo.Id, waitUse.ID)
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

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err := service.InsertOrUpdateAnswer(qa.Qid, qa.Aid, qa.WithCard, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

func CreateQa(c *gin.Context) {
	var param request.CreateQaParams
	_ = c.ShouldBindJSON(&param)

	ApiVerify := utils.Rules{
		"Qid": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(param, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	question := model.ZhihuQuestion{
		Qid:   param.Qid,
		Title: param.Qname,
		Type:  param.Type,
	}

	err := service.CreateQuestion(question, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建问题失败，%v", err), c)
	}

	if param.Aid != "" {
		err = service.InsertOrUpdateAnswer(param.Qid, param.Aid, param.WithCard, waitUse.ID)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("创建回答失败，%v", err), c)
		}
	}

	response.OkWithMessage("创建问答成功", c)
}

func CreateQuestion(c *gin.Context) {
	var param request.CreateQuestionParams
	_ = c.ShouldBindJSON(&param)

	ApiVerify := utils.Rules{
		"Qid": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(param, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	question := model.ZhihuQuestion{
		Qid:  param.Qid,
		Type: param.Type,
	}

	err := service.CreateQuestion(question, param.UserID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建问题失败，%v", err), c)
	}

	if param.Aid != "" {
		err = service.InsertOrUpdateAnswer(param.Qid, param.Aid, param.WithCard, param.UserID)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("创建回答失败，%v", err), c)
		}
	}

	response.OkWithMessage("创建问答成功", c)
}

func DeleteQa(c *gin.Context) {
	var a model.ZhihuQuestionSub
	_ = c.ShouldBindJSON(&a)

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	a.UserID = waitUse.ID

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

func GetQaByQid(c *gin.Context) {
	var param request.GetQaByQidParams
	_ = c.ShouldBindJSON(&param)

	err, qa := service.GetQaByQid(param.Qid, param.UserID)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.ZhihuQaResponse{Qa: qa}, c)
	}
}
