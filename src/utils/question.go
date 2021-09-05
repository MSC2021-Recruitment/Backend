package utils

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
)

type QuestionNotFoundError struct{}

func (m *QuestionNotFoundError) Error() string {
	return "question not found"
}

func IsQuestionExists(req *models.Question) bool {
	err := global.DATABASE.First(req).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
