package api

import (
	"MSC2021/src/models"
	"MSC2021/src/models/auth"
	"MSC2021/src/models/requests"
	response "MSC2021/src/models/responses"
	"MSC2021/src/services"
	"MSC2021/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var req requests.LoginRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	var user *models.User
	if utils.VerifyEmailFormat(req.Account) {
		user, err = services.LoginWithEmail(&models.User{
			Password: req.Password,
			Email:    req.Account,
		})
	} else if utils.VerifyMobileFormat(req.Account) {
		user, err = services.LoginWithTel(&models.User{
			Password:  req.Password,
			Telephone: req.Account,
		})
	}
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	if user == nil {
		response.FailWithMessage("Account not found.", ctx)
		ctx.Abort()
		return
	}
	token := utils.NewToken()
	tokenStr, err := token.CreateToken(auth.TokenClaims{
		UserID:         user.ID,
		Admin:          user.Admin,
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
	}
	_ = services.PutTokenInWhitelist(user.ID, tokenStr)
	ctx.Header("New-Token", tokenStr)
	response.Ok(ctx)
}

func RegisterHandler(ctx *gin.Context) {
	var req requests.RegisterRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	_, err = services.RegisterWithUser(models.User{
		Name:      req.Name,
		Password:  req.Password,
		StudentID: req.StudentID,
		Admin:     false,
		Major:     req.Major,
		Telephone: req.Telephone,
		Email:     req.Email,
		QQ:        req.QQ,
		Level:     req.Level,
		Wanted:    req.Wanted,
		Intro:     req.Intro,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.Ok(ctx)
}

func GetProfileHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		response.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	user, err := services.GetUserProfile(claims.UserID)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.OkWithData(user, ctx)
}

func ChangeProfileHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		response.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	var req requests.ChangeProfileRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	err = services.ChangeUserProfile(models.User{
		ID:        claims.UserID,
		Name:      req.Name,
		StudentID: req.StudentID,
		Major:     req.Major,
		Telephone: req.Telephone,
		Email:     req.Email,
		QQ:        req.QQ,
		Level:     req.Level,
		Wanted:    req.Wanted,
		Intro:     req.Intro,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.Ok(ctx)
}

func ChangePasswordHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		response.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	var req requests.ChangePasswordRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	_, err = services.ChangePassword(claims.UserID, req.Password, req.NewPassword)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.Ok(ctx)
}

func LogoutHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		response.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	err := services.ExpireToken(claims.UserID)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.Ok(ctx)
}
