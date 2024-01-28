package dietController

import (
	"encoding/json"
	"india-fit/domain/service/dietService"
	"log"
	"net/http"
	"strconv"
)

func ListDietOfUser(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("id")
	dietId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		log.Println("error in converting dietId into int64", id)
	}

	myRes, myErr := dietService.ListDiet(dietId)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
