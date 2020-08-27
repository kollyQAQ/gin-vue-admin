package response

import "gin-vue-admin/model"

type ZhihuCollectionResponse struct {
	Collection model.ZhihuCollection `json:"collection"`
}

type ZhihuCollectionListResponse struct {
	CollectionList []model.ZhihuCollection `json:"collection_list"`
}
