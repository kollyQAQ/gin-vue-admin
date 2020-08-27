package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitZhihuRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("zhihu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("getGoodsList", v1.GetGoodsList)       // 获取商品列表
		ApiRouter.POST("getQaList", v1.GetQuestionAnswer)     // 获取商品列表
		ApiRouter.POST("getCollectionList", v1.GetCollection) // 获取收藏列表
	}
}
