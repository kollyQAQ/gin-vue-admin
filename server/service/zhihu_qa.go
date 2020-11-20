package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

func GetZhihuQuestionAnswer(qa model.ZhihuQuestionAnswer, info request.PageInfo, order string,
	desc bool, answer, withCard string, userID uint) (err error, list interface{}, total int) {
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

	if answer == "0" {
		db = db.Where("aid = ''")
	}

	if answer == "1" {
		db = db.Where("aid != ''")
	}

	if withCard != "" {
		db = db.Where("with_card = ?", withCard)
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
			err = db.Debug().Select("id,aid,answer_total,like_num,qid,qname,rank,sub_id,three_day_add_view,today_add_answer,today_add_view,type,view_total,with_card,date_format(update_time, '%Y-%m-%d %H:%i') as update_time").
				Order(OrderStr, true).Find(&qaList).Error
		} else {
			err = db.Debug().Select("id,aid,answer_total,like_num,qid,qname,rank,sub_id,three_day_add_view,today_add_answer,today_add_view,type,view_total,with_card,date_format(update_time, '%Y-%m-%d %H:%i') as update_time").
				Order("today_add_view desc", true).Find(&qaList).Error
		}
	}
	return err, qaList, total
}

func GetQaById(id string, userID uint) (err error, qa model.ZhihuQuestionAnswer) {
	err = global.GVA_DB.Where("qid = ?", id).Where("user_id = ?", userID).First(&qa).Error
	return
}

func GetQaByQid(qid string, userID int) (err error, qa model.ZhihuQuestionAnswer) {
	err = global.GVA_DB.Where("qid = ?", qid).Where("user_id = ?", userID).First(&qa).Error
	return
}

func InsertOrUpdateAnswer(qid, aid, withCard string, userID uint) (err error) {
	find := model.ZhihuAnswer{Qid: qid, UserID: userID}
	update := model.ZhihuAnswer{Aid: aid, WithCard: withCard}
	insert := model.ZhihuAnswer{Qid: qid, UserID: userID, Aid: aid, WithCard: withCard}

	//	Where("qid = ?", answer.Qid).Update("aid", answer.Aid)
	//global.GVA_DB.Model(&model.ZhihuAnswer{}).

	return global.GVA_DB.Where(find).Assign(update).FirstOrCreate(&insert).Error
}

func CreateQuestion(question model.ZhihuQuestion, qtype int, userID uint) (err error) {
	findOne := global.GVA_DB.Where("qid = ?", question.Qid).Find(&model.ZhihuQuestion{}).Error
	if findOne != nil {
		err = global.GVA_DB.Create(&question).Error
	}

	sub := &model.ZhihuQuestionSub{
		UserID: userID,
		Qid:    question.Qid,
		Type:   qtype,
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
