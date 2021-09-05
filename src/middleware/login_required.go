package middleware

import (
	"MSC2021/src/global"
	"MSC2021/src/models/auth"
	"MSC2021/src/models/responses"
	"MSC2021/src/services"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenStr string
		if token := ctx.Request.Header.Get("Authentication"); token != "" {
			global.LOGGER.Warn("Auth failed, no authentication found")
			response.FailWithDetailed(gin.H{"reload": true}, "Token is Null!", ctx)
			ctx.Abort()
			return
		} else {
			tokenStr = token
		}
		j := NewToken()
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		whitelistToken, err := services.GetTokenInWhitelist(claims.UserID)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "Token is not in whitelist.", ctx)
			ctx.Abort()
		}
		if whitelistToken != tokenStr {
			response.FailWithDetailed(gin.H{"reload": true}, "Token is not same with the one in whitelist.", ctx)
			ctx.Abort()
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

type Token struct {
	SigningKey []byte
}

func NewToken() *Token {
	return &Token{
		[]byte(global.CONFIG.JWT.SigningKey),
	}
}

func (j *Token) CreateToken(claims auth.TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *Token) ParseToken(token string) (*auth.TokenClaims, error) {
	res, err := jwt.ParseWithClaims(token, &auth.TokenClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if res != nil {
		if claims, ok := res.Claims.(*auth.TokenClaims); ok && res.Valid {
			return claims, nil
		}
		return nil, errors.New("token is not valid or claims broken")
	}
	return nil, errors.New("token is invalid")
}
