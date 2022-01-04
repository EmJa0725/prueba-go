package main

import (
	"api_proyect/parte1/api"
	"api_proyect/parte1/config"
	"api_proyect/parte1/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	//Api routes
	apiRouter.HandleFunc("/days-report", api.DaysReport).Methods("GET")
	//Middlewares
	router.Use(middlewares.LoggingMiddleware)

	fmt.Println("Server running on port " + config.PORT)
	http.ListenAndServe(":"+config.PORT, router)

}
