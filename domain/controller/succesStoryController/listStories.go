package succesStoryController

import (
	"encoding/json"
	"india-fit/domain/service/succesStoryService"
	"log"
	"net/http"
	"strconv"
)

func ListStories(res http.ResponseWriter, req *http.Request) {

	count := req.URL.Query().Get("count")
	converted_count, err := strconv.ParseInt(count, 10, 64)
	if err == nil {
		log.Println("error in converting count into int64", count)
	}

	myRes, myErr := succesStoryService.ListStories(converted_count)
	if myErr != nil {
		log.Println("error in generating response")
		return
	}

	json.NewEncoder(res).Encode(myRes)
	return
}
