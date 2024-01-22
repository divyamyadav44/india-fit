package userService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/userRepo"
	"log"
)

func CreateUser(newUser *model.NewUser) (user model.User, err error) {

	//Password hashing
	user, err = userRepo.CreateUser(newUser)
	if err != nil {
		log.Println("error in creating user: ", err)
		return
	}
	return user, nil
}
