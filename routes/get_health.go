package routes

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/uelker/go-monitoring/metrics"
	"net/http"
)

func GetHealth(uptime *metrics.Uptime) func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		uptime, uptimeUnit := uptime.Value()
		healthResponse := HealthResponse{Uptime: UptimeResponse{Value: uptime, Unit: uptimeUnit}}

		jsonResponse, err := json.Marshal(healthResponse)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonResponse)
	}
}
