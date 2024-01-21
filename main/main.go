package main

import (
	"fmt"
	"india-fit/main/intialize"
	"india-fit/main/route"
	"log"
	"net/http"
)

func main() {
	intialize.Init()

	err := http.ListenAndServe(":3333", route.HttpHandler)
	if err != nil {
		log.Println("err in creating handler")
	}
	fmt.Println("Welcome to india-fit")

}
