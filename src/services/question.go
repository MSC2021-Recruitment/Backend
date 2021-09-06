package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/utils"
)

func AdminCreateQuestion(creatorId uint, title string, group string, content string) error {

	ques := models.Question{
		CreatorID: creatorId,
		Title:     title,
		Group:     group,
		Content:   content,
	}
	err := global.DATABASE.Create(&ques).Error
	return err
}

type AdminQuestionRes struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	Group           string `json:"group"`
	SubmissionCount int    `json:"submission-count"`
}

func GetQuestionList() ([]models.Question, error) {
	var ques []models.Question
	err := global.DATABASE.Select([]string{"id", "title", "group"}).Find(&ques).Error
	return ques, err
}

func AdminGetQuestionList() ([]AdminQuestionRes, error) {
	var ques []models.Question
	var res []AdminQuestionRes
	err := global.DATABASE.Select([]string{"id", "title", "group"}).Find(&ques).Error
	if err != nil {
		return res, err
	}
	for _, q := range ques {
		answeredUserCount := global.DATABASE.Model(&q).Association("AnsweredUsers").Count()
		res = append(res, AdminQuestionRes{
			ID:              q.ID,
			Title:           q.Title,
			SubmissionCount: int(answeredUserCount),
			Group:           q.Group,
		})
	}
	return res, nil
}

type AdminQuestionDetailRes struct {
	ID      uint          `json:"id"`
	Creator models.User   `json:"creator"`
	Title   string        `json:"title"`
	Group   string        `json:"group"`
	Content string        `json:"content"`
	Users   []models.User `json:"users"`
}

type QuestionDetailRes struct {
	ID      uint        `json:"id"`
	Creator models.User `json:"creator"`
	Title   string      `json:"title"`
	Group   string      `json:"group"`
	Content string      `json:"content"`
}

func GetQuestionDetail(questionId uint) (QuestionDetailRes, error) {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(&ques).Error
	if err != nil {
		global.LOGGER.Sugar().Warnf("GET Question Failed: %s", err)
		return QuestionDetailRes{}, err
	}
	creator := models.User{ID: ques.CreatorID}
	err = global.DATABASE.First(&creator).Error
	if err != nil {
		global.LOGGER.Sugar().Warnf("GET Question Creator %d Failed: %s", ques.CreatorID, err)
		creator = models.User{}
	}
	return QuestionDetailRes{
		ID:      ques.ID,
		Creator: creator,
		Title:   ques.Title,
		Group:   ques.Group,
		Content: ques.Content,
	}, nil
}

func AdminGetQuestionDetail(questionId uint) (AdminQuestionDetailRes, error) {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(&ques).Error
	if err != nil {
		return AdminQuestionDetailRes{}, err
	}
	creator := models.User{ID: ques.CreatorID}
	err = global.DATABASE.First(&creator).Error
	if err != nil {
		return AdminQuestionDetailRes{}, err
	}
	var users []models.User
	err = global.DATABASE.Model(&ques).Association("AnsweredUsers").Find(&users)
	if err != nil {
		return AdminQuestionDetailRes{}, err
	}
	return AdminQuestionDetailRes{
		ID:      ques.ID,
		Creator: creator,
		Title:   ques.Title,
		Group:   ques.Group,
		Content: ques.Content,
		Users:   users,
	}, nil
}

func AdminChangeQuestionDetail(questionId uint, title string, group string, content string) error {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.First(&ques).Error
	if err != nil {
		return err
	}
	ques.Title = title
	ques.Group = group
	ques.Content = content
	err = global.DATABASE.Save(&ques).Error
	return err
}

func AdminDeleteQuestion(questionId uint) error {
	ques := models.Question{ID: questionId}
	err := global.DATABASE.Delete(&ques).Error
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
			Content: ans,
		}
		err = global.DATABASE.Create(&submission).Error
		if err != nil {
			return err
		}
		err = global.DATABASE.Model(&ques).Association("AnsweredUsers").Append(&user)
		if err != nil {
			return err
		}
		err = global.DATABASE.Model(&user).Association("AnsweredQuestions").Append(&ques)
		if err != nil {
			return err
		}
		err = global.DATABASE.Model(&ques).Association("Submissions").Append(&submission)
		if err != nil {
			return err
		}
		err = global.DATABASE.Model(&user).Association("Submissions").Append(&submission)
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
	err := global.DATABASE.Where("question_id = ? AND user_id = ?", questionId, userId).Find(&submissions).Error
	if err != nil {
		global.LOGGER.Sugar().Warnf("Get user%d's submission of %d failed: %s", userId, questionId, err.Error())
	}
	return submissions, err
}
