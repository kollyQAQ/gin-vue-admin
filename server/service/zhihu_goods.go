package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

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

func GetGoodsById(id float64) (err error, goods model.ZhihuGoodsWithContent) {
	err = global.GVA_DB.Where("id = ?", id).First(&goods).Error
	return
}

func UpdateGoods(goods model.ZhihuGoodsWithContent) (err error) {
	var oldGoods model.ZhihuGoodsWithContent

	err = global.GVA_DB.Where("id = ?", goods.ID).First(&oldGoods).Error

	if oldGoods.SkuID != goods.SkuID {
		flag := global.GVA_DB.Where("sku_id = ?", goods.SkuID).Find(&model.SysApi{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同sku_id")
		}
	}

	err = global.GVA_DB.Save(&goods).Error

	return err
}
