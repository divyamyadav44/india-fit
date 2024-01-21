package userService

import (
	"india-fit/domain/model"
)

func GetUser(id int64) (user model.User, err error) {
	user = model.User{
		Id:           id,
		Name:         "Divyam",
		Email:        "divyamyadav44@gmail.com",
		MobileNumber: "5634737374",
	}

	return user, nil
}
