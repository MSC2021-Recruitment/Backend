package utils

import (
	"MSC2021/src/global"
	"MSC2021/src/models/auth"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

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
