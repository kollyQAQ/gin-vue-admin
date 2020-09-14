package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
)

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

func GetQaById(id string) (err error, qa model.ZhihuQuestionAnswer) {
	err = global.GVA_DB.Where("qid = ?", id).First(&qa).Error
	return
}

func UpdateQuestion(question model.ZhihuQuestion) (err error) {
	var old model.ZhihuQuestion

	err = global.GVA_DB.Where("qid = ?", question.Qid).First(&old).Error

	if old.Qid != question.Qid {
		flag := global.GVA_DB.Where("qid = ?", question.Qid).Find(&model.ZhihuQuestion{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同qid")
		}
	}

	err = global.GVA_DB.Save(&question).Error

	return err
}

func UpdateAnswer(qid, aid, with_card string) (err error) {
	find := model.ZhihuAnswer{Qid: qid}
	update := model.ZhihuAnswer{Aid: aid, WithCard: with_card}
	insert := model.ZhihuAnswer{Qid: qid, Aid: aid, WithCard: with_card}

	//	Where("qid = ?", answer.Qid).Update("aid", answer.Aid)
	//global.GVA_DB.Model(&model.ZhihuAnswer{}).

	return global.GVA_DB.Where(find).Assign(update).FirstOrCreate(&insert).Error

}

func CreateQuestion(question model.ZhihuQuestion) (err error) {
	findOne := global.GVA_DB.Where("qid = ?", question.Qid).Find(&model.ZhihuQuestion{}).Error
	if findOne == nil {
		return errors.New("存在相同qid")
	} else {
		err = global.GVA_DB.Create(&question).Error
	}
	return err
}

func DeleteQuestion(question model.ZhihuQuestion) (err error) {
	err = global.GVA_DB.Delete(question).Error
	return err
}

func QueryQuestionHistory(qid string) (err error, list []model.ZhihuQuestionHistory) {
	db := global.GVA_DB.Model(&model.ZhihuQuestionHistory{})
	var qaList []model.ZhihuQuestionHistory

	err = db.Where("qid = ?", qid).Find(&qaList).Error

	return err, qaList
}

func QueryQaStat() (err error, list []*model.ZhihuQaStat) {
	querySql := `
		SELECT 
			a.type, a.total, b.answer, c.card_answer
		FROM (
			SELECT 
				type,
				count(*) AS total
			FROM view_question_answer
			GROUP BY type
		) a INNER JOIN (
			SELECT 
				type,
				count(*) AS answer
			FROM view_question_answer
			WHERE aid != ''
			GROUP BY type
		) b on a.type = b.type
		LEFT JOIN (
			SELECT 
				type,
				count(*) AS card_answer
			FROM view_question_answer
			WHERE aid != '' and with_card = 1
			GROUP BY type
		) c on a.type = c.type
	`
	err = global.GVA_DB.Raw(querySql).Scan(&list).Error

	return err, list
}

func QueryStat() (err error, data resp.ZhihuStat) {
	querySql := `
		SELECT 
			a.no_fee_goods_num,
			b.today_fee,
			c.todo_number,
			d.today_view
		FROM (
			SELECT '1' as id, COUNT(*) as no_fee_goods_num FROM t_zhihu_goods where fee_rate = 0
		) a
		LEFT JOIN (
			SELECT '1' as id, ROUND(SUM(fee)) as today_fee FROM t_zhihu_order where to_days(order_time) = to_days(now())
		) b ON a.id = b.id
		LEFT JOIN (
			SELECT '1' as id, COUNT(*) as todo_number FROM t_todo where status = 0
		) c ON a.id = c.id
		LEFT JOIN (
			SELECT '1' as id, SUM(today_add_view) as today_view FROM view_question_answer
		) d ON a.id = d.id
	`
	err = global.GVA_DB.Raw(querySql).Scan(&data).Error

	return err, data
}

func GetQuestionTypeMap() map[int]string {
	return map[int]string{
		1: "礼物",
		2: "乳胶枕",
		3: "按摩仪",
		4: "教师节礼物",
		5: "生日礼物",
	}
}
