package authSerivce

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"india-fit/domain/model"
	"india-fit/domain/repo/authRepo"
	"india-fit/domain/util"
	"log"
	"time"
)

func Login(newlogin *model.UserLogin) (authTokens model.AuthTokens, err error) {

	//Password hashing
	hashedPassword, err := HashPassword(newlogin.Password)
	if err != nil {
		log.Println("error in generating hashed password: ", err)
		return
	}
	hashedPasswordFromDb, err := authRepo.Login(newlogin.Password, newlogin.Password)
	if err != nil {
		log.Println("error in getting password from repo layer: ", err)
		return
	}

	passErr := bcrypt.CompareHashAndPassword(hashedPasswordFromDb, hashedPassword)
	if passErr != nil {
		log.Println("error in password matching from repo layer: ", passErr)
		err = errors.New("Incorrect Password")
		//return
	}

	userClaims := model.UserClaims{
		Id:   "1",
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	authTokens.AccessToken, err = util.NewAccessToken(userClaims)
	if err != nil {
		log.Println("error in generating token", err)
		return
	}

	refreshClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	}

	authTokens.RefreshToken, err = util.NewRefreshToken(refreshClaims)
	if err != nil {
		log.Fatal("error creating refresh token")
	}

	return
}

func HashPassword(password string) (bytes []byte, err error) {
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err

}
