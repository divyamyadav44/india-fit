package util

import (
	"github.com/golang-jwt/jwt"
	"india-fit/domain/model"
	"os"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func NewAccessToken(claims model.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("sampleSecretKey")))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
