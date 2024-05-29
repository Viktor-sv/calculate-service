package main

import (
	"calculate-service/pkg/server/handler/calculate"
	"calculate-service/pkg/server/handler/middleware"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const calculate = "calculate"

func main() {

	fmt.Print("hello")
	router := httprouter.New()
	router.POST("/"+calculate, middleware.Validate(hanlder.Calculate))
	log.Fatal(http.ListenAndServe("localhost:8989", router))
}
