package middleware

import (
	"MSC2021/src/models/auth"
	"MSC2021/src/models/responses"

	"github.com/gin-gonic/gin"
)

func AdminRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimsRaw, exists := ctx.Get("claims")
		if !exists {
			responses.FailWithDetailed(gin.H{"reload": true}, "Not login yet.", ctx)
			ctx.Abort()
			return
		}
		claims := claimsRaw.(auth.TokenClaims)
		if !claims.Admin {
			responses.FailWithDetailed(gin.H{"reload": true}, "Not Admin!", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
