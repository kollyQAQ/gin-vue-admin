package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
)

func GetZhihuQuestionAnswer(qa model.ZhihuQuestionAnswer, info request.PageInfo, order string,
	desc bool, userID uint) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ZhihuQuestionAnswer{})
	var qaList []model.ZhihuQuestionAnswer

	db = db.Where("user_id = ?", userID)

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
			err = db.Debug().Order(OrderStr, true).Find(&qaList).Error
		} else {
			err = db.Debug().Order("today_add_view desc", true).Find(&qaList).Error
		}
	}
	return err, qaList, total
}

func GetQaById(id string) (err error, qa model.ZhihuQuestionAnswer) {
	err = global.GVA_DB.Where("qid = ?", id).First(&qa).Error
	return
}

func UpdateAnswer(qid, aid, withCard string, userID uint) (err error) {
	find := model.ZhihuAnswer{Qid: qid, UserID: userID}
	update := model.ZhihuAnswer{Aid: aid, WithCard: withCard}
	insert := model.ZhihuAnswer{Qid: qid, UserID: userID, Aid: aid, WithCard: withCard}

	//	Where("qid = ?", answer.Qid).Update("aid", answer.Aid)
	//global.GVA_DB.Model(&model.ZhihuAnswer{}).

	return global.GVA_DB.Where(find).Assign(update).FirstOrCreate(&insert).Error

}

func CreateQuestion(question model.ZhihuQuestion, userID uint) (err error) {
	findOne := global.GVA_DB.Where("qid = ?", question.Qid).Find(&model.ZhihuQuestion{}).Error
	if findOne != nil {
		err = global.GVA_DB.Create(&question).Error
	}

	sub := &model.ZhihuQuestionSub{
		UserID: userID,
		Qid:    question.Qid,
	}

	findOne = global.GVA_DB.Where("user_id = ?", sub.UserID).
		Where("qid = ?", sub.Qid).
		Find(&model.ZhihuQuestionSub{}).Error

	if findOne == nil {
		return errors.New("该问题已经被你添加过了")
	} else {
		err = global.GVA_DB.Create(&sub).Error
	}

	return err
}

func DeleteQuestion(question model.ZhihuQuestionSub) (err error) {
	err = global.GVA_DB.Delete(question).Error
	// TODO：如果这个问题没有其他人关注，应该把问题本身也删掉, 问题的回答也删掉
	return err
}

func QueryQuestionHistory(qid string) (err error, list []model.ZhihuQuestionHistory) {
	db := global.GVA_DB.Model(&model.ZhihuQuestionHistory{})
	var qaList []model.ZhihuQuestionHistory

	err = db.Where("qid = ?", qid).Find(&qaList).Error

	return err, qaList
}

func QueryQaStat(userID uint) (err error, list []*model.ZhihuQaStat) {
	querySql := `
		SELECT 
			a.type, a.total, b.answer, c.card_answer
		FROM (
			SELECT 
				type,
				count(*) AS total
			FROM view_question_answer
			WHERE user_id = ?
			GROUP BY type
		) a INNER JOIN (
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

func GetQuestionTypeMap() map[int]string {
	return map[int]string{
		0: "未分类",
		1: "礼物",
		2: "乳胶枕",
		3: "按摩仪",
		4: "教师节礼物",
		5: "生日礼物",
	}
}
