package api

import (
	"MSC2021/src/models/auth"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartInterviewHandler(ctx *gin.Context) {
	err := services.StartInterview()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

func StopInterviewHandler(ctx *gin.Context) {
	err := services.StopInterview()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

func GetInterviewStatusHandler(ctx *gin.Context) {
	status, err := services.GetInterviewStatus()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.OkWithData(status, ctx)
}

func UserSignInInterviewHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		responses.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	userId := claims.UserID
	err := services.UserSignInInterview(userId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

func AdminGetSuggestInterviewUserHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	user, err := services.AdminGetSuggestInterviewUser(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.OkWithData(user, ctx)
}

func AdminInterviewHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	user, err := services.AdminGetNextInterviewUser(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.OkWithData(user, ctx)
}

func AdminGetInterviewContentHandler(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	userId, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("Status in wrong format.", ctx)
		ctx.Abort()
		return
	}
	content, err := services.AdminGetInterviewContent(uint(groupId), uint(userId))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.OkWithData(content, ctx)
}

func AdminUpdateInterviewContentHandler(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	userId, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("Status in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.AdminUpdateInterviewContent(uint(groupId), uint(userId), ctx.PostForm("content"))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

func AdminDeleteInterviewContentHandler(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("groupId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	userId, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("Status in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.AdminDeleteInterviewContent(uint(groupId), uint(userId))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}
