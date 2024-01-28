package userController

import (
	"encoding/json"
	"fmt"
	"india-fit/domain/model"
	"india-fit/domain/service/userService"
	"india-fit/domain/util"
	"log"
	"net/http"
)

func UpdateUser(res http.ResponseWriter, req *http.Request) {

	var user model.User
	err := util.ReadRequestBody(req.Body, &user)
	if err != nil {
		http.Error(res, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	myRes, myErr := userService.UpdateUser(&user)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
