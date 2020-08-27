package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int
func GetZhihuGoodsList(goods model.ZhihuGoods, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuGoods{})
	var goodsList []model.ZhihuGoods

	if goods.SkuID != "" {
		db = db.Where("sku_id LIKE ?", "%"+goods.SkuID+"%")
	}

	if goods.SkuName != "" {
		db = db.Where("sku_name LIKE ?", "%"+goods.SkuName+"%")
	}

	if goods.ThirdCategory != "" {
		db = db.Where("cid3_name LIKE ?", "%"+goods.ThirdCategory+"%")
	}

	if goods.FeeRate > 0 {
		db = db.Where("fee_rate <= ?", goods.FeeRate)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, goodsList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr, true).Find(&goodsList).Error
		} else {
			err = db.Order("fee_rate desc", true).Find(&goodsList).Error
		}
	}
	return err, goodsList, total
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int
func GetZhihuQuestionAnswer(qa model.ZhihuQuestionAnswer, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuQuestionAnswer{})
	var qaList []model.ZhihuQuestionAnswer

	if qa.Qid != "" {
		db = db.Where("qid LIKE ?", "%"+qa.Qid+"%")
	}

	if qa.Type > 0 {
		db = db.Where("type = ?", qa.Type)
	}

	if qa.Qname != "" {
		db = db.Where("qname LIKE ?", "%"+qa.Qname+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, qaList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr, true).Find(&qaList).Error
		} else {
			err = db.Order("today_add_view desc", true).Find(&qaList).Error
		}
	}
	return err, qaList, total
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int
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
			err = db.Order("type", true).Find(&collectionList).Error
		}
	}
	return err, collectionList, total
}
