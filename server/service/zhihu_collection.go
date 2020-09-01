package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

func GetZhihuCollectionList(c model.ZhihuCollection, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuCollection{})
	var collectionList []model.ZhihuCollection

	if c.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+c.Type+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, collectionList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr, true).Find(&collectionList).Error
		} else {
			err = db.Order("priority desc, type", true).Find(&collectionList).Error
		}
	}
	return err, collectionList, total
}

func GetCollectionById(id float64) (err error, c model.ZhihuCollection) {
	err = global.GVA_DB.Where("id = ?", id).First(&c).Error
	return
}

func UpdateCollection(c model.ZhihuCollection) (err error) {
	err = global.GVA_DB.Save(&c).Error

	return err
}

func CreateCollection(c model.ZhihuCollection) (err error) {
	err = global.GVA_DB.Create(&c).Error
	return err
}

func DeleteCollection(c model.ZhihuCollection) (err error) {
	err = global.GVA_DB.Delete(c).Error
	return err
}
