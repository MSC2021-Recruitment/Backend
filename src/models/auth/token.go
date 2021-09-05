package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	UserID uint `json:"user-id"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}
