package response

import "gin-vue-admin/model"

type ZhihuCollectionResponse struct {
	Collection model.ZhihuCollectionWithContent `json:"collection"`
}

type ZhihuCollectionListResponse struct {
	CollectionList []model.ZhihuCollection `json:"collection_list"`
}
