package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/uelker/go-monitoring/metrics"
	"github.com/uelker/go-monitoring/routes"
	"net/http"
	"time"
)

func main() {
	uptime := metrics.NewUptime(time.Now())

	router := httprouter.New()

	router.GET("/", routes.GetRoot())
	router.GET("/health", routes.GetHealth(uptime))

	server := http.Server{Addr: ":3000", Handler: router}
	err := server.ListenAndServe()

	log.Fatal().Err(err).Msg("server failed to start")
}
