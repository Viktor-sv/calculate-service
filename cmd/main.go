package main

import (
	"calculate-service/pkg/server/handler/calculate"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const calculate = "calculate"

func main() {

	fmt.Print("hello")
	router := httprouter.New()
	//TODO add middleware
	router.POST("/"+calculate, hanlder.Calculate)

	log.Fatal(http.ListenAndServe("localhost:8989", router))
}
