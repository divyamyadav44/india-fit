package main

import (
	"fmt"
	"india-fit/main/intialize"
	"log"
	"net/http"
)

func main() {
	intialize.Init()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("err in creating handler")
	}

	fmt.Println("Welcome to india-fit")

}
