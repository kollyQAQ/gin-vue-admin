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

	if goods.SkuName != "" {
		db = db.Where("sku_name LIKE ?", "%"+goods.SkuName+"%")
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
			err = db.Order("id", true).Find(&goodsList).Error
		}
	}
	return err, goodsList, total
}
