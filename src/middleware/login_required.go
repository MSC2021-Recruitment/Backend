package middleware

import (
	"MSC2021/src/global"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"MSC2021/src/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenStr string
		if token := ctx.Request.Header.Get("Authentication"); token != "" {
			global.LOGGER.Warn("Auth failed, no authentication found")
			response.FailWithDetailed(gin.H{"login": false}, "Token is Null!", ctx)
			ctx.Abort()
			return
		} else {
			tokenStr = token
		}
		j := utils.NewToken()
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			response.FailWithDetailed(gin.H{"login": false}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		whitelistToken, err := services.GetTokenInWhitelist(claims.UserID)
		if err != nil {
			response.FailWithDetailed(gin.H{"login": false}, "Token is not in whitelist.", ctx)
			ctx.Abort()
			return
		}
		if whitelistToken != tokenStr {
			response.FailWithDetailed(gin.H{"login": false}, "Token is not same with the one in whitelist.", ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < global.CONFIG.JWT.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("New-Token", newToken)
			_ = services.PutTokenInWhitelist(newClaims.UserID, newToken)
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
