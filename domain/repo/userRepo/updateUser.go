package userRepo

import "india-fit/domain/model"

func UpdateUser(newUser *model.User) (user model.User, err error) {

	user = model.User{
		Id:           1,
		Name:         newUser.Name,
		Email:        newUser.Email,
		MobileNumber: newUser.MobileNumber,
		DOB:          newUser.DOB,
	}
	return
}
