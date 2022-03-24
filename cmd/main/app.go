package main

import (
	"github.com/aakosarev/rest-api/internal/user"
	"github.com/aakosarev/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
