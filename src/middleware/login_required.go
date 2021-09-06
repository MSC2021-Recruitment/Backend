package middleware

import (
	"MSC2021/src/global"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"MSC2021/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenStr string
		if token := ctx.Request.Header.Get("Authorization"); token == "" {
			global.LOGGER.Warn("Auth failed, no authentication found")
			responses.FailWithDetailed(gin.H{"login": false}, "Token is Null!", ctx)
			ctx.Abort()
			return
		} else {
			tokenStr = strings.ReplaceAll(token, "Bearer ", "")
		}
		j := utils.NewToken()
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			responses.FailWithDetailed(gin.H{"login": false}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		whitelistToken, err := services.GetTokenInWhitelist(claims.UserID)
		if err != nil {
			responses.FailWithDetailed(gin.H{"login": false}, "Token is not in whitelist.", ctx)
			ctx.Abort()
			return
		}
		if whitelistToken != tokenStr {
			responses.FailWithDetailed(gin.H{"login": false}, "Token is not same with the one in whitelist.", ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < global.CONFIG.JWT.BufferTime {
			global.LOGGER.Sugar().Debugf("token is going to be expired: %s -> %s", time.Now().String(), time.Unix(claims.ExpiresAt, 0).String())
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("New-Token", newToken)
			_ = services.PutTokenInWhitelist(newClaims.UserID, newToken)
		}
		ctx.Set("claims", *claims)
		ctx.Next()
	}
}
