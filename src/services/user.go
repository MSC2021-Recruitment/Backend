package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/utils"
	"errors"
	"gorm.io/gorm"
)

func RegisterWithUser(req models.User) (res models.User, err error) {
	var user models.User
	if !errors.Is(global.DATABASE.Where("telephone = ?", req.Telephone).
		First(&user).Error, gorm.ErrRecordNotFound) &&
		!errors.Is(global.DATABASE.Where("email = ?", req.Email).
			First(&user).Error, gorm.ErrRecordNotFound) {
		return res, errors.New("user is already registered")
	}
	req.Password, _ = utils.HashPassword(req.Password)
	err = global.DATABASE.Create(&req).Error
	global.LOGGER.Sugar().Infof("New User Registered: %s %s %s", req.StudentID, req.Name, req.Wanted)
	return req, err
}

func LoginWithTel(req *models.User) (res *models.User, err error) {
	var user models.User
	err = global.DATABASE.Where("telephone = ?", req.Telephone).First(&user).Error
	if err == nil {
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			err = errors.New("password is wrong")
		}
	}
	return &user, err
}

func LoginWithEmail(req *models.User) (res *models.User, err error) {
	var user models.User
	err = global.DATABASE.Where("email = ?", req.Email).First(&user).Error
	if err == nil {
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			err = errors.New("password is wrong")
		}
	}
	return &user, err
}

func ChangePassword(id uint, oldPassword string, newPassword string) (res *models.User, err error) {
	return ChangePasswordWithUser(&models.User{
		ID:       id,
		Password: oldPassword,
	}, newPassword)
}

func ChangePasswordWithUser(req *models.User, newPassword string) (userInter *models.User, err error) {

	var user = models.User{
		ID: req.ID,
	}
	err = global.DATABASE.First(&user).Error
	if err == nil {
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			err = errors.New("password is not true")
			return nil, err
		} else {
			newPassword, _ = utils.HashPassword(newPassword)
			err = global.DATABASE.First(&user).Update("password", newPassword).Error
		}
	}
	err = ExpireToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func GetUserList() ([]models.User, error) {
	var users []models.User
	result := global.DATABASE.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserProfile(userId uint) (models.User, error) {
	user := models.User{ID: userId}
	err := global.DATABASE.First(&user)
	if err.Error != nil {
		return models.User{}, err.Error
	}
	return user, nil
}

func DeleteUser(userId uint) error {
	user := models.User{ID: userId}
	err := global.DATABASE.Delete(&user).Error
	err = ExpireToken(userId)
	if err != nil {
		return err
	}
	return err
}

func ChangeUserProfile(req models.User) error {
	err := global.DATABASE.Model(&models.User{}).
		Select("*").
		Omit("password").
		Omit("create_at").
		Updates(&req).Error
	return err
}
