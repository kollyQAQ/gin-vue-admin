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

func GetGoodsList(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchGoodsParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, list, total := service.GetZhihuGoodsList(sp.ZhihuGoods, sp.PageInfo, sp.OrderKey, sp.Desc, waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取商品数据失败，%v", err), c)
	} else {
		err, category := service.GetZhihuGoodsCategoryList(waitUse.ID)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("获取类目数据失败，%v", err), c)
		} else {
			response.OkWithData(resp.PageResult{
				List:     list,
				Category: category,
				Total:    total,
				Page:     sp.PageInfo.Page,
				PageSize: sp.PageInfo.PageSize,
			}, c)
		}
	}
}

func GetGoodsCategory(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	err, category := service.GetZhihuGoodsCategoryList(waitUse.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取类目数据失败，%v", err), c)
	} else {
		response.OkWithData(category, c)
	}
}

func GetGoodsByID(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}

	err, goods := service.GetGoodsByID(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.ZhihuGoodsResponse{Goods: goods}, c)
	}
}

func GetGoodsBySkuID(c *gin.Context) {
	var idInfo request.GetByIdStr
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}

	err, goods := service.GetGoodsBySkuID(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(goods, c)
	}
}

func UpdateGoods(c *gin.Context) {
	var goods model.ZhihuGoodsWithContent
	_ = c.ShouldBindJSON(&goods)
	ApiVerify := utils.Rules{
		"SkuID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(goods, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	err := service.UpdateGoods(goods)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

func CreateGoods(c *gin.Context) {
	var goods model.ZhihuGoods
	_ = c.ShouldBindJSON(&goods)
	ApiVerify := utils.Rules{
		"SkuID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(goods, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	goods.UserID = waitUse.ID

	err := service.CreateGoods(goods)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteGoods(c *gin.Context) {
	var a model.ZhihuGoods
	_ = c.ShouldBindJSON(&a)
	//ApiVerify := utils.Rules{
	//	"ID": {utils.NotEmpty()},
	//}
	//ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	//if ApiVerifyErr != nil {
	//	response.FailWithMessage(ApiVerifyErr.Error(), c)
	//	return
	//}
	err := service.DeleteGoods(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
