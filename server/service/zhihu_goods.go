package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
)

func GetZhihuGoodsList(goods model.ZhihuGoods, info request.PageInfo, order string,
	desc bool, userID uint) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuGoods{})
	var goodsList []model.ZhihuGoods

	db = db.Where("user_id = ?", userID)

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

	err = db.Debug().Count(&total).Error

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
			err = db.Debug().Select("id, sku_id, sku_name, price, fee_rate, fee, jd_sale, cid3_name").
				Order(OrderStr, true).Find(&goodsList).Error
		} else {
			err = db.Debug().Select("id, sku_id, sku_name, price, fee_rate, fee, jd_sale, cid3_name").
				Order("fee_rate desc", true).Find(&goodsList).Error
		}
	}
	return err, goodsList, total
}

func GetGoodsById(id float64) (err error, goods model.ZhihuGoodsWithContent) {
	err = global.GVA_DB.Debug().Where("id = ?", id).First(&goods).Error
	return
}

func UpdateGoods(goods model.ZhihuGoodsWithContent) (err error) {
	var oldGoods model.ZhihuGoodsWithContent

	err = global.GVA_DB.Where("id = ?", goods.ID).First(&oldGoods).Error

	if oldGoods.SkuID != goods.SkuID {
		flag := global.GVA_DB.Where("sku_id = ?", goods.SkuID).Find(&model.ZhihuGoodsWithContent{}).RecordNotFound()
		if !flag {
			return errors.New("该商品已经存在")
		}
	}

	err = global.GVA_DB.Save(&goods).Error

	return err
}

func CreateGoods(goods model.ZhihuGoods) (err error) {

	findOne := global.GVA_DB.Where("sku_id = ?", goods.SkuID).
		Where("user_id = ?", goods.UserID).
		Find(&model.ZhihuGoods{}).Error

	if findOne == nil {
		return errors.New("该商品已经被你添加过了")
	} else {
		goods.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		err = global.GVA_DB.Create(&goods).Error
	}
	return err
}

func DeleteGoods(goods model.ZhihuGoods) (err error) {
	err = global.GVA_DB.Delete(goods).Error
	return err
}
