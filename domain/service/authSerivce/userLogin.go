package authSerivce

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"india-fit/domain/model"
	"india-fit/domain/repo/authRepo"
	"log"
)

func Login(newlogin *model.UserLogin) (jwtToken string, err error) {

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

	log.Println("hashedPassword--->", string(hashedPassword), newlogin.Password)

	passErr := bcrypt.CompareHashAndPassword(hashedPasswordFromDb, hashedPassword)
	if passErr != nil {
		log.Println("error in password matching from repo layer: ", passErr)
		err = errors.New("Incorrect Password")
		return
	}
	jwtToken = "success"
	return
}

func HashPassword(password string) (bytes []byte, err error) {
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err

}
