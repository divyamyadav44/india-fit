package route

import (
	"github.com/gorilla/mux"
	"india-fit/domain/controller/userController"
	"net/http"
)

var HttpHandler http.Handler

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("api/v1/getUser", userController.GetUser).Methods(http.MethodGet)

	HttpHandler = r

}
