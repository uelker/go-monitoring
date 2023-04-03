package main

import (
	"net/http"
)

func main() {
	server := http.Server{Addr: ":3000"}
	server.ListenAndServe()
}
