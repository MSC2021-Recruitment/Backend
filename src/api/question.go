package api

import (
	"MSC2021/src/models"
	"MSC2021/src/models/auth"
	"MSC2021/src/models/requests"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQuestionListHandler(ctx *gin.Context) {
	questions, err := services.GetQuestionList()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(questions, ctx)
}

func GetQuestionDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	question, err := services.GetQuestionDetail(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(question, ctx)
}

func AnswerQuestionHandler(ctx *gin.Context) {
	var req requests.AnswerQuestionRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		responses.FailWithMessage("Data is invalid.", ctx)
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
	quesId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.AnswerQuestion(models.User{ID: userId}, models.Question{ID: uint(quesId)}, req.Answer)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminGetQuestionListHandler(ctx *gin.Context) {
	questions, err := services.AdminGetQuestionList()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	if questions == nil {
		responses.FailWithMessage("No Questions Here.", ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(questions, ctx)
}

func AdminGetQuestionDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	question, err := services.AdminGetQuestionDetail(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(question, ctx)
}

func AdminChangeQuestionDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	var req requests.UpdateQuestionRequest
	err = ctx.ShouldBind(&req)
	if err != nil {
		responses.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	err = services.AdminChangeQuestionDetail(uint(id), req.Title, req.Group, req.Content)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminDeleteQuestionHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.AdminDeleteQuestion(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminCreateQuestionHandler(ctx *gin.Context) {
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		responses.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	userId := claims.UserID
	var req requests.CreateQuestionRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		responses.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	err = services.AdminCreateQuestion(userId, req.Title, req.Group, req.Content)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}
