package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
)

func GetLogList(c model.ZhihuLog, info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuLog{})
	var collectionList []model.ZhihuLog

	err = db.Count(&total).Error

	if c.Type != "" {
		db = db.Where("`type` = ?", c.Type)
	}

	if err != nil {
		return err, collectionList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Debug().Select("id, type, event, msg, update_time").
			Order("id desc", true).Find(&collectionList).Error
	}

	return err, collectionList, total
}

func QueryQaStat(userID uint) (err error, list []*model.ZhihuQaStat) {
	querySql := `
		SELECT 
			a.type, 
			a.type_desc, 
			a.total, IFNULL(b.answer,0) as answer, 
			IFNULL(c.card_answer,0) as card_answer
		FROM (
			SELECT 
				type,
				IFNULL(typ.label, '未分类') as type_desc,
				count(*) AS total
			FROM view_question_answer ans
			LEFT JOIN t_zhihu_type typ ON ans.type = typ.id
			WHERE ans.user_id = ?
			GROUP BY type
		) a LEFT JOIN (
			SELECT 
				type,
				count(*) AS answer
			FROM view_question_answer
			WHERE aid != '' and user_id = ?
			GROUP BY type
		) b on a.type = b.type
		LEFT JOIN (
			SELECT 
				type,
				count(*) AS card_answer
			FROM view_question_answer
			WHERE aid != '' and with_card = 1 and user_id = ?
			GROUP BY type
		) c on a.type = c.type
	`
	err = global.GVA_DB.Raw(querySql, userID, userID, userID).Scan(&list).Error

	return err, list
}

func QueryStat(userID uint) (err error, data resp.ZhihuStat) {
	querySql := `
		SELECT 
			a.no_fee_goods_num,
			b.today_fee,
			c.todo_number,
			d.today_view
		FROM (
			SELECT '1' as id, COUNT(*) as no_fee_goods_num FROM t_zhihu_goods where fee_rate = 0 and user_id = ?
		) a
		LEFT JOIN (
			SELECT '1' as id, ifnull(ROUND(SUM(fee)),0) as today_fee FROM t_zhihu_order where to_days(order_time) = to_days(now()) and user_id = ?
		) b ON a.id = b.id
		LEFT JOIN (
			SELECT '1' as id, COUNT(*) as todo_number FROM t_todo where status = 0 and user_id = ?
		) c ON a.id = c.id
		LEFT JOIN (
			SELECT '1' as id, ifnull(SUM(today_add_view),0) as today_view FROM view_question_answer where user_id = ?
		) d ON a.id = d.id
	`
	err = global.GVA_DB.Debug().Raw(querySql, userID, userID, userID, userID).Scan(&data).Error

	return err, data
}

func QuerySearchStat(userID uint) (err error, list []*resp.ZhihuSearchStat) {
	querySql := `
		SELECT article_id, remark, rank, update_time
		FROM t_zhihu_search
		WHERE user_id = ?;
	`
	err = global.GVA_DB.Debug().Raw(querySql, userID).Scan(&list).Error

	return err, list
}
