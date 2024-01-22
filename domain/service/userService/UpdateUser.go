package userService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/userRepo"
	"log"
)

func UpdateUser(user *model.User) (updateUser model.User, err error) {

	updateUser, err = userRepo.UpdateUser(user)
	if err != nil {
		log.Println("error in creating user: ", err)
		return
	}
	return updateUser, nil
}
