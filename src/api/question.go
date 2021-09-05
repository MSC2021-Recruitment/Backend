package api

import (
	"MSC2021/src/models"
	"MSC2021/src/models/auth"
	"MSC2021/src/models/requests"
	response "MSC2021/src/models/responses"
	"MSC2021/src/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQuestionListHandler(ctx *gin.Context) {
	questions, err := services.GetQuestionList()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	if questions == nil {
		response.FailWithMessage("No Questions Here.", ctx)
		ctx.Abort()
		return
	}
	response.OkWithData(questions, ctx)
}

func GetQuestionDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	question, err := services.GetQuestionDetail(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.OkWithData(question, ctx)
}

func AnswerQuestionHandler(ctx *gin.Context) {
	var req requests.AnswerQuestionRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage("Data is invalid.", ctx)
		ctx.Abort()
		return
	}
	claimsRaw, exists := ctx.Get("claims")
	if !exists {
		response.FailWithMessage("Not login yet.", ctx)
		ctx.Abort()
		return
	}
	claims := claimsRaw.(auth.TokenClaims)
	userId := claims.UserID
	quesId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.AnswerQuestion(models.User{ID: userId}, models.Question{ID: uint(quesId)}, req.Answer)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	response.Ok(ctx)
}

func AdminGetQuestionListHandler(ctx *gin.Context) {

}

func AdminGetQuestionDetailHandler(ctx *gin.Context) {

}

func AdminChangeQuestionDetailHandler(ctx *gin.Context) {

}

func AdminDeleteQuestionHandler(ctx *gin.Context) {

}

func AdminCreateQuestionHandler(ctx *gin.Context) {

}
