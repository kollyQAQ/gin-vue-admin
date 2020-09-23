package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitZhihuRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("zhihu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("getGoodsList", v1.GetGoodsList)         // 获取商品列表
		ApiRouter.POST("getGoodsCategory", v1.GetGoodsCategory) // 获取商品类目
		ApiRouter.POST("getGoodsById", v1.GetGoodsById)         // 获取单个商品信息
		ApiRouter.POST("createGoods", v1.CreateGoods)           // 创建新商品
		ApiRouter.POST("updateGoods", v1.UpdateGoods)           // 修改商品信息
		ApiRouter.POST("deleteGoods", v1.DeleteGoods)           // 删除商品

		ApiRouter.POST("getQaList", v1.GetQuestionAnswer) // 获取问答列表
		ApiRouter.POST("getQaById", v1.GetQaById)         // 获取单个问答信息
		ApiRouter.POST("createQa", v1.CreateQa)           // 创建新问答
		ApiRouter.POST("updateQa", v1.UpdateQa)           // 修改问答信息
		ApiRouter.POST("deleteQa", v1.DeleteQa)           // 删除问答
		ApiRouter.POST("queryQaHistory", v1.QueryHistory) // 查看问答历史

		ApiRouter.POST("getCollectionList", v1.GetCollection)     // 获取收藏列表
		ApiRouter.POST("getCollectionListV2", v1.GetCollectionV2) // 获取收藏列表
		ApiRouter.POST("getCollectionById", v1.GetCollectionById) // 更新单个收藏信息
		ApiRouter.POST("createCollection", v1.CreateCollection)   // 创建新收藏
		ApiRouter.POST("updateCollection", v1.UpdateCollection)   // 修改收藏信息
		ApiRouter.POST("deleteCollection", v1.DeleteCollection)   // 删除收藏

		ApiRouter.POST("queryStat", v1.QueryStat) // 统计数据
		ApiRouter.POST("queryLog", v1.QueryLog)   // 日志数据

		ApiRouter.POST("queryMusicList", v1.GetMusicList) // 查询Music列表

		ApiRouter.POST("queryPlanList", v1.GetPlanList) // 查询TODO列表

		ApiRouter.POST("queryTypeList", v1.GetTypeList) // 查询 type 列表

		ApiRouter.POST("queryTodoList", v1.GetTodoList) // 查询TODO列表
		ApiRouter.POST("addTodo", v1.AddTodo)           // 添加TODO
		ApiRouter.POST("updateTodo", v1.UpdateTodo)     // 更新TODO
		ApiRouter.POST("deleteTodo", v1.DeleteTodo)     // 删除TODO
	}
}
