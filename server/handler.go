package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/definition"
	"net/http"
)

type httpHandlerStruct struct {
	weatherForecast *definition.IWeatherForecast
}

var httpHandler httpHandlerStruct

func initHttpHandler(weatherForecast *definition.IWeatherForecast) {
	if weatherForecast == nil {
		logrus.Fatal("can not init httpHandler - weatherForecast is nil")
	}
	httpHandler = httpHandlerStruct{
		weatherForecast: weatherForecast,
	}
}

func (h *httpHandlerStruct) GetWeatherInsight(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	result, status, err := (*h.weatherForecast).GetWeatherInsight(params["condition"], params["lat"], params["lon"])
	if err != nil {
		h.handleError(err, w, status)
		return
	}
	var response []byte
	if len(result) == 0 {
		response, _ = json.Marshal("no results were found")

	} else {
		response, _ = json.Marshal(result)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *httpHandlerStruct) handleError(err error, w http.ResponseWriter, status int) {
	logrus.WithError(err).Error()
	w.WriteHeader(status)
	response, _ := json.Marshal(err.Error())
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return
}
