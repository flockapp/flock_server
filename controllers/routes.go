package controllers

import (
	"github.com/flockapp/flock_server/controllers/api"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/v0/api").Subrouter()
	apiRouter.HandleFunc("/", api.V0_Get_API).Methods("GET")

	return router
}

func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
