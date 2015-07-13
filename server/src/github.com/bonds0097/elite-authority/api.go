package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	FIElD_LENGTH_MIN int = 3
	FIELD_LENGTH_MAX int = 255
)

type ApiError struct {
	Err  string
	Code int
}

func setApiHandlers() {
	pathPrefix := "/api/"
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc(pathPrefix+"commodities/", handleErrors(getCommodities)).Methods("GET")
	router.HandleFunc(pathPrefix+"commodities/", handleErrors(addCommodity)).Methods("POST")
	router.HandleFunc(pathPrefix+"commodities/list/", handleErrors(listCommodities)).Methods("GET")
	router.HandleFunc(pathPrefix+"stations/", handleErrors(getRecentStations)).Methods("GET")
	router.HandleFunc(pathPrefix+"stations/{slug}", handleErrors(getStationDetails)).Methods("GET")
	router.HandleFunc(pathPrefix+"stations/", handleErrors(addStation)).Methods("POST")
	// router.HandleFunc(pathPrefix+"stations/{slug}", handleErrors(updateStation)).Methods("PUT")
	router.HandleFunc(pathPrefix+"systems/search/", handleErrors(findSystems)).Methods("POST")
	router.HandleFunc(pathPrefix+"snapshots/", handleErrors(addSnapshot)).Methods("POST")

	http.Handle(pathPrefix, router)
}

func handleErrors(f func(w http.ResponseWriter, r *http.Request) *ApiError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiErr := f(w, r)
		if apiErr == nil {
			return
		}

		log.Println("Error handling request from: " + r.RemoteAddr + "\nRequest: " + r.Method + " " +
			r.RequestURI + "\n" + apiErr.Err)
		http.Error(w, apiErr.Err, apiErr.Code)
	}
}
