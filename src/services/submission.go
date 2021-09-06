package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/utils"
)

func GetSubmissionOfQuestion(questionId uint, userId uint) ([]models.Submission, error) {
	var err error
	var submissions []models.Submission
	err = global.DATABASE.Model(&models.Submission{}).Where("user_id = ? AND question_id = ?", userId, questionId).First(&submissions).Error
	if err != nil {
		global.LOGGER.Sugar().Warnf("Get user%d's submission of %d failed: %s", userId, questionId, err.Error())
	}
	return submissions, err
}

func GetSubmittedUserOfQuestion(questionId uint) ([]models.User, error) {
	var err error
	if !utils.IsQuestionExists(&models.Question{ID: questionId}) {
		err = &utils.QuestionNotFoundError{}
	}
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = global.DATABASE.Model(&models.Question{ID: questionId}).Association("AnsweredUsers").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
