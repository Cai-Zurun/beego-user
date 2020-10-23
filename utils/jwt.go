package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	expireDuration = time.Hour * 24
	scretKey = "thisisakey"
)

type LoginClaims struct {
	userEmail string
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		userEmail:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	return token.SignedString([]byte(scretKey))
}