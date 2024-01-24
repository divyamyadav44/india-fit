package userService

import (
	"golang.org/x/crypto/bcrypt"
	"india-fit/domain/model"
	"india-fit/domain/repo/userRepo"
	"log"
)

func CreateUser(newUser *model.NewUser) (user model.User, err error) {

	//Password hashing
	newUser.HashedPassword, err = HashPassword(newUser.Password)
	if err != nil {
		log.Println("error in generating hashed password: ", err)
		return
	}
	user, err = userRepo.CreateUser(newUser)
	if err != nil {
		log.Println("error in creating user: ", err)
		return
	}
	return user, nil
}

func HashPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}
