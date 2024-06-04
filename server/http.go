package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/config"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/definition"
	"net/http"
)

var httpServer *http.Server

func StartHTTP(weatherForecast *definition.IWeatherForecast) *http.Server {
	router := mux.NewRouter()
	initHttpHandler(weatherForecast)
	registerRoutes(router)
	httpServer = &http.Server{
		Addr:    config.Static.HTTPServerPort,
		Handler: router,
	}
	go listenAndServe(httpServer)
	return httpServer
}

func listenAndServe(server *http.Server) {
	logrus.Infof("Starting http server on addr %v", config.Static.HTTPServerPort)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("failed to start http server")
	}
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/weather/insight", httpHandler.GetWeatherInsight).Methods("GET")
}

func Shutdown() {
	if httpServer != nil {
		if err := httpServer.Shutdown(context.Background()); err != nil {
			logrus.Fatalf("Server shutdown failed: %v", err)
		}
	}
}
