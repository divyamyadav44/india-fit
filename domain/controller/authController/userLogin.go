package authController

import (
	"encoding/json"
	"fmt"
	"india-fit/domain/model"
	"india-fit/domain/service/authSerivce"
	"india-fit/domain/util"
	"log"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {

	var userLoginDetails model.UserLogin
	err := util.ReadRequestBody(req.Body, &userLoginDetails)
	if err != nil {
		http.Error(res, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	myRes, myErr := authSerivce.Login(&userLoginDetails)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
