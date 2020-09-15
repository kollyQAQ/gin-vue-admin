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

func GetCollection(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchCollectionParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, list, total := service.GetZhihuCollectionList(sp.ZhihuCollection, sp.PageInfo, sp.OrderKey, sp.Desc, waitUse.ID)
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

func GetCollectionById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, collection := service.GetCollectionById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.ZhihuCollectionResponse{Collection: collection}, c)
	}
}

func UpdateCollection(c *gin.Context) {
	var collection model.ZhihuCollectionWithContent
	_ = c.ShouldBindJSON(&collection)
	ApiVerify := utils.Rules{
		"SkuID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(collection, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.UpdateCollection(collection)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

func CreateCollection(c *gin.Context) {
	var collection model.ZhihuCollectionWithContent
	_ = c.ShouldBindJSON(&collection)

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	collection.UserID = waitUse.ID

	err := service.CreateCollection(collection)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteCollection(c *gin.Context) {
	var a model.ZhihuCollection
	_ = c.ShouldBindJSON(&a)
	//ApiVerify := utils.Rules{
	//	"ID": {utils.NotEmpty()},
	//}
	//ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	//if ApiVerifyErr != nil {
	//	response.FailWithMessage(ApiVerifyErr.Error(), c)
	//	return
	//}
	err := service.DeleteCollection(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
