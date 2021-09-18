package api

import (
	"MSC2021/src/models/auth"
	"MSC2021/src/models/requests"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGroupListHandler(ctx *gin.Context) {
	uid := ctx.DefaultQuery("id", "-1")
	userId, err := strconv.ParseInt(uid, 10, 32)
	if err != nil {
		return
	}
	var ans []services.GroupRes
	ans, err = services.GetGroupList(int(userId))

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(ans, ctx)
}

func JoinGroupHandler(ctx *gin.Context) {
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
	err = services.JoinGroup(uint(id), userId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func LeaveGroupHandler(ctx *gin.Context) {
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
	err = services.LeaveGroup(uint(id), userId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminGetGroupListHandler(ctx *gin.Context) {
	ans, err := services.AdminGetGroupList()
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(ans, ctx)
}

func AdminGetGroupDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	ans, err := services.AdminGetGroupDetail(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.OkWithData(ans, ctx)
}

func AdminChangeGroupDetailHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	var req requests.ChangeGroupDetailRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	err = services.ChangeGroupDetail(uint(id), req.Name, req.Description)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminDeleteGroupHandler(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		responses.FailWithMessage("ID in wrong format.", ctx)
		ctx.Abort()
		return
	}
	err = services.DeleteGroup(uint(id))
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}

func AdminCreateGroupHandler(ctx *gin.Context) {
	var req requests.CreateGroupRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	err = services.CreateGroup(req.Name, req.Description)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		ctx.Abort()
		return
	}
	responses.Ok(ctx)
}
