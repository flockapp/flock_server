package controllers

import (
	"github.com/flockapp/flock_server/controllers/api"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/v0/api").Subrouter()
	apiRouter.HandleFunc("/", api.API_Get_API).Methods("GET")
	apiRouter.HandleFunc("/verify", Use(api.API_Get_API, RequireUserToken)).Methods("GET")
	apiRouter.HandleFunc("/types", Use(api.API_Get_Types, RequireUserToken)).Methods("GET")
	apiRouter.HandleFunc("/events", Use(api.API_Get_Events, RequireUserToken)).Methods("GET")
	apiRouter.HandleFunc("/events", Use(api.API_Create_Event, RequireUserToken)).Methods("POST")
	apiRouter.HandleFunc("/events/{eventId}", Use(api.API_Get_Event_Details, RequireUserToken)).Methods("GET")
	apiRouter.HandleFunc("/events/guests", Use(api.API_Get_Guest_Events, RequireUserToken)).Methods("GET")
	apiRouter.HandleFunc("/events/guests", Use(api.API_Put_Guest_Into_Event, RequireUserToken)).Methods("PUT")
	apiRouter.HandleFunc("/events/{eventId}/guests", Use(api.API_Get_Guests_From_Event, RequireUserToken)).Methods("GET")

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", AUTH_Post_Login).Methods("POST")
	authRouter.HandleFunc("/register", AUTH_Post_Register).Methods("POST")
	return router
}

func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
