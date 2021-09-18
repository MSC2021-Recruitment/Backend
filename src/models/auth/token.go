package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	UserID uint   `json:"user-id"`
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}
