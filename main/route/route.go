package route

import (
	"github.com/gorilla/mux"
	"india-fit/domain/controller/authController"
	"india-fit/domain/controller/dietController"
	"india-fit/domain/controller/succesStoryController"
	"india-fit/domain/controller/userController"
	"net/http"
)

var HttpHandler http.Handler

func InitRouter() {
	r := mux.NewRouter()

	//Auth
	r.HandleFunc("/api/v1/login", authController.Login).Methods(http.MethodPost)

	// User Module
	r.HandleFunc("/api/v1/createUser", userController.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/getUser", userController.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/updateUser", userController.UpdateUser).Methods(http.MethodPut)

	//Diet Module
	r.HandleFunc("/api/v1/listDiet", dietController.ListDietOfUser).Methods(http.MethodGet)

	//Success Story Module
	r.HandleFunc("/api/v1/succesStory/getStoryById", succesStoryController.GetStory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/succesStory/ListStories", succesStoryController.ListStories).Methods(http.MethodGet)

	HttpHandler = r

}
