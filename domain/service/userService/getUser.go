package userService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/userRepo"
	"log"
)

func GetUser(id int64) (user model.User, err error) {

	user, err = userRepo.GetUser(id)
	if err != nil {
		log.Println("error in getting user details", err)
		return
	}
	return user, err
}
