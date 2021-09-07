package api

import (
	"MSC2021/src/models/responses"
	"MSC2021/src/services"

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
