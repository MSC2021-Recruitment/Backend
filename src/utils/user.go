package utils

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
)

type UserNotFoundError struct{}

func (m *UserNotFoundError) Error() string {
	return "user not found"
}

func IfUserExists(req *models.User) bool {
	err := global.DATABASE.First(&req).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
