package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"MSC2021/src/models/requests"
	"MSC2021/src/utils"
	"errors"
	"gorm.io/gorm"
)

func Register(request requests.RegisterRequest) (res models.User, err error) {
	return RegisterWithUser(models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Major:     request.Major,
		Telephone: request.Telephone,
		QQ:        request.QQ,
		Level:     request.Level,
		Wanted:    request.Wanted,
		Intro:     request.Intro,
	})
}

func RegisterWithUser(req models.User) (res models.User, err error) {
	var user models.User
	if !errors.Is(global.DATABASE.Where("name = ?", req.Name).First(&user).Error, gorm.ErrRecordNotFound) &&
		!errors.Is(global.DATABASE.Where("email = ?", req.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return res, errors.New("user is already registered")
	}
	req.Password, _ = utils.HashPassword(req.Password)
	req.Admin = false
	err = global.DATABASE.Create(&req).Error
	global.LOGGER.Info("New User Registered:")
	return req, err
}

func Login(loginForm requests.LoginRequest) (res *models.User, err error) {
	if utils.VerifyEmailFormat(loginForm.Account) {
		return LoginWithEmail(&models.User{
			Email:    loginForm.Account,
			Password: loginForm.Password,
		})
	} else {
		return LoginWithTel(&models.User{
			Telephone: loginForm.Account,
			Password:  loginForm.Password,
		})
	}
}

func LoginWithTel(req *models.User) (res *models.User, err error) {
	var user models.User
	err = global.DATABASE.Where("telephone = ?", req.Name).First(&user).Error
	if err == nil {
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			err = errors.New("password is not true")
		}
	}
	return &user, err
}

func LoginWithEmail(req *models.User) (res *models.User, err error) {
	var user models.User
	err = global.DATABASE.Where("email = ?", req.Email).First(&user).Error
	if err == nil {
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			err = errors.New("password is not true")
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
		if !utils.CheckPasswordHash(user.Password, req.Password) {
			err = errors.New("password is not true")
		} else {
			newPassword, _ = utils.HashPassword(newPassword)
			err = global.DATABASE.First(&user).Update("password", newPassword).Error
		}
	}
	return &user, err
}
