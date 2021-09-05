package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/utils"
)

func GetQuestionList() ([]models.Question, error) {
	var ques []models.Question
	err := global.DATABASE.Select([]string{"title", "group"}).Find(&ques).Error
	return ques, err
}

func GetQuestionDetail(questionId uint) (models.Question, error) {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(ques).Error
	return ques, err
}

func AnswerQuestion(user models.User, ques models.Question, ans string) error {
	var err error
	if !utils.IfUserExists(&user) {
		err = &utils.UserNotFoundError{}
	}
	if !utils.IsQuestionExists(&ques) && err == nil {
		err = &utils.QuestionNotFoundError{}
	}

	if err == nil {
		submission := models.Submission{
			Submitter: user,
			Question:  ques,
			Content:   ans,
		}
		err = global.DATABASE.Create(submission).Error
	}
	if err != nil {
		global.LOGGER.Sugar().Warnf("%s %s answered question %s, %s occured: %s", user.StudentID, user.Name, ques.Title, err.Error(), ans)
	} else {
		global.LOGGER.Sugar().Infof("%s %s answered question %s.", user.StudentID, user.Name, ques.Title)
	}
	return err
}
