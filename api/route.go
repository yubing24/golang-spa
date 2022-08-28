package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewApiServer(conf Config) *mux.Router {
	router := mux.NewRouter()
	router.Methods(http.MethodGet).Handler(helloWorld(conf))
	return router
}
