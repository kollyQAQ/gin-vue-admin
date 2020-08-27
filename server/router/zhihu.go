package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitZhihuRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("zhihu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.GET("getGoodsList", v1.GetGoodsList)       // 获取商品列表
		ApiRouter.GET("getQaList", v1.GetQuestionAnswer)     // 获取商品列表
		ApiRouter.GET("getCollectionList", v1.GetCollection) // 获取收藏列表
	}
}
