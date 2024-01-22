package route

import (
	"github.com/gorilla/mux"
	"india-fit/domain/controller/succesStoryController"
	"india-fit/domain/controller/userController"
	"net/http"
)

var HttpHandler http.Handler

func InitRouter() {
	r := mux.NewRouter()

	// User Module
	r.HandleFunc("/api/v1/createUser", userController.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/getUser", userController.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/updateUser", userController.UpdateUser).Methods(http.MethodPut)

	//Success Story Module
	r.HandleFunc("/api/v1/succesStory/getStoryById", succesStoryController.GetStory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/succesStory/ListStories", succesStoryController.ListStories).Methods(http.MethodGet)

	HttpHandler = r

}
