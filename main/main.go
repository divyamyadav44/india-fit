package main

import (
	"india-fit/main/intialize"
	"india-fit/main/route"
	"log"
	"net/http"
)

func main() {
	intialize.Init()

	log.Println("service is up and running on port 3333")
	err := http.ListenAndServe(":3333", route.HttpHandler)
	if err != nil {
		log.Println("err in creating handler")
	}
}
