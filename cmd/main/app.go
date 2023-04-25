package main

import (
	"net"
	"net/http"
	"time"

	"github.com/RAprogramm/rest-api-tutorial/internal/user"
	"github.com/RAprogramm/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	logger.Info("starting server...")
	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port 1234")
	logger.Fatal(server.Serve(listener))
}
