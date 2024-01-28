package model

import "github.com/golang-jwt/jwt"

type UserLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
