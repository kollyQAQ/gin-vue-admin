package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
)

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetGoodsList(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchGoodsParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetZhihuGoodsList(sp.ZhihuGoods, sp.PageInfo, sp.OrderKey, sp.Desc)
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

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
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

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetCollection(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchCollectionParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetZhihuCollectionList(sp.ZhihuCollection, sp.PageInfo, sp.OrderKey, sp.Desc)
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
