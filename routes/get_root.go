package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetRoot() func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Ok"))
	}
}
