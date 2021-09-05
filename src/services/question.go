package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/utils"
)

func AdminCreateQuestion(creatorId uint, title string, group string, content string) error {
	creator := models.User{ID: creatorId}

	ques := models.Question{
		Creator: creator,
		Title:   title,
		Group:   group,
		Content: content,
	}
	err := global.DATABASE.Create(ques).Error
	return err
}

type AdminQuestionRes struct {
	Title           string `json:"title"`
	Group           string `json:"group"`
	SubmissionCount int    `json:"submission-count"`
}

func GetQuestionList() ([]models.Question, error) {
	var ques []models.Question
	err := global.DATABASE.Select([]string{"title", "group"}).Find(&ques).Error
	return ques, err
}

func AdminGetQuestionList() ([]AdminQuestionRes, error) {
	var ques []models.Question
	var res []AdminQuestionRes
	err := global.DATABASE.Select([]string{"title", "group"}).Find(&ques).Error
	if err != nil {
		return res, err
	}
	for _, q := range ques {
		var submissionCount int64
		err = global.DATABASE.Model(&models.Submission{}).Where("QuestionRefer = ?", q.ID).Count(&submissionCount).Error
		if err != nil {
			return res, err
		}
		res = append(res, AdminQuestionRes{
			Title:           q.Title,
			SubmissionCount: int(submissionCount),
			Group:           q.Group,
		})
	}
	return res, nil
}

type AdminQuestionDetailRes struct {
	Title string        `json:"title"`
	Group string        `json:"group"`
	Users []models.User `json:"users"`
}

func GetQuestionDetail(questionId uint) (models.Question, error) {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(ques).Error
	return ques, err
}

func AdminGetQuestionDetail(questionId uint) (AdminQuestionDetailRes, error) {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(ques).Error
	if err != nil {
		return AdminQuestionDetailRes{}, err
	}
	var users []models.User
	err = global.DATABASE.Model(&models.Submission{}).Where("QuestionRefer = ?", ques.ID).Association("Submitter").Find(&users)
	if err != nil {
		return AdminQuestionDetailRes{}, err
	}
	return AdminQuestionDetailRes{
		Title: ques.Title,
		Group: ques.Group,
		Users: users,
	}, nil
}

func AdminChangeQuestionDetail(questionId uint, title string, group string, content string) error {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(ques).Error
	if err != nil {
		return err
	}
	ques.Title = title
	ques.Group = group
	ques.Content = content
	err = global.DATABASE.Save(ques).Error
	return err
}

func AdminDeleteQuestion(questionId uint) error {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.Delete(ques).Error
	return err
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

func AdminGetSubmissionOfQuestionAndUser(questionId uint, userId uint) ([]models.Submission, error) {
	var submissions []models.Submission
	err := global.DATABASE.Model(&models.Submission{}).Where("QuestionRefer = ? AND SubmitterRefer = ?", questionId, userId).Find(&submissions).Error
	return submissions, err
}
