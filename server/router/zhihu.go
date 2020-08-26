package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitZhihuRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("zhihu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("getGoodsList", v1.GetGoodsList) // 获取商品列表
	}
}
