package userController

import (
	"encoding/json"
	"india-fit/domain/service/userService"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		log.Println("error in converting id into int64", userId)
	}

	myRes, myErr := userService.GetUser(userId)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
