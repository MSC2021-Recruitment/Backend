package api

import (
	"MSC2021/src/models"
	"MSC2021/src/models/requests"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminGetUserListHandler(ctx *gin.Context) {
	users, err := services.GetUserList()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(users, ctx)
}

func AdminGetUserProfileHandler(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("User ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	user, err := services.GetUserProfile(uint(userId))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(user, ctx)
}

func AdminChangeUserProfileHandler(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("User ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	var userReq requests.ChangeUserProfileRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	user := models.User{
		ID:        uint(userId),
		Name:      userReq.Name,
		Email:     userReq.Email,
		StudentID: userReq.StudentID,
		Major:     userReq.Major,
		Admin:     userReq.Admin,
		Telephone: userReq.Telephone,
		QQ:        userReq.QQ,
		Level:     userReq.Level,
		Intro:     userReq.Intro,
	}
	err = services.ChangeUserProfile(user)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	_ = services.ExpireToken(user.ID)
	responses.Ok(ctx)
}

func AdminDeleteUserHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("User ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.DeleteUser(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminChangeUserPasswordHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("User ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	var userReq requests.ChangeUserPasswordRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	_, err = services.ChangePassword(uint(id), userReq.Password, userReq.NewPassword)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}
