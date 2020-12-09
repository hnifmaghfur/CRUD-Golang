package main

import (
	"github.com/gorilla/mux"
)

func server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Active).Methods("GET")
	router.HandleFunc("/api/user", getUser).Methods("GET")
	router.HandleFunc("/api/user", postUser).Methods("POST")
	router.HandleFunc("/api/user", patchUser).Methods("PATCH")
	return router
}


