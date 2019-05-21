package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/working_day/{date}", CalculateWorkingDays).Methods("GET")
	return router
}
