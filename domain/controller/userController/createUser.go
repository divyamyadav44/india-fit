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

func CreateUser(res http.ResponseWriter, req *http.Request) {

	var newUser model.NewUser
	err := util.ReadRequestBody(req.Body, &newUser)
	if err != nil {
		http.Error(res, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	myRes, myErr := userService.CreateUser(&newUser)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
