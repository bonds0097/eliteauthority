package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ApiError struct {
	Err  string
	Code int
}

func setApiHandlers() {
	pathPrefix := "/api/"
	router := mux.NewRouter()
	router.HandleFunc(pathPrefix+"commodities/", handleErrors(listCommodities)).Methods("GET")
	router.HandleFunc(pathPrefix+"commodities/", handleErrors(addCommodity)).Methods("POST")
	router.HandleFunc(pathPrefix+"snapshots/", handleErrors(addSnapshot)).Methods("POST")

	http.Handle(pathPrefix, router)
}

func handleErrors(f func(w http.ResponseWriter, r *http.Request) *ApiError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiErr := f(w, r)
		if apiErr == nil {
			return
		}

		log.Println(apiErr.Err)
		http.Error(w, apiErr.Err, apiErr.Code)
	}
}
