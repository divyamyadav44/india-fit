package succesStoryController

import (
	"encoding/json"
	"india-fit/domain/service/succesStoryService"
	"log"
	"net/http"
	"strconv"
)

func GetStory(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("id")
	converted_id, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		log.Println("error in converting id into int64", id)
	}

	myRes, myErr := succesStoryService.GetStory(converted_id)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
