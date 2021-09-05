package api

import (
	"MSC2021/src/models/auth"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSubmissionOfQuestionHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		responses.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	userId := claims.UserID
	submissions, err := services.GetSubmissionOfQuestion(uint(id), userId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(submissions, ctx)
}

func AdminGetSubmittedUserOfQuestionHandler(ctx *gin.Context) {
	quesId, err := strconv.ParseUint(ctx.Param("quesId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	users, err := services.GetSubmittedUserOfQuestion(uint(quesId))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(users, ctx)
}

func AdminGetSubmissionOfQuestionAndUserHandler(ctx *gin.Context) {
	quesId, err := strconv.ParseUint(ctx.Param("quesId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	userId, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		responses.FailWithMessage("User ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	submissions, err := services.AdminGetSubmissionOfQuestionAndUser(uint(quesId), uint(userId))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(submissions, ctx)
}
