package main

import (
	"calculate-service/pkg/server/handler/calculate"
	"calculate-service/pkg/server/handler/middleware"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const calculate = "calculate"

func main() {
	shutdownChannel := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	router := httprouter.New()
	router.POST("/"+calculate, middleware.Validate(hanlder.Calculate))

	s := &http.Server{
		Addr:        ":8989",
		Handler:     router,
		IdleTimeout: 60 * time.Second,
	}

	// server start
	go func() {
		fmt.Printf("sing: %v", <-signals)
		close(shutdownChannel)
	}()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-shutdownChannel
	_ = s.Shutdown(context.Background())
}
