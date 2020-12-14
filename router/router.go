package router

import (
	"github.com/gorilla/mux"
	"github.com/hnifmaghfur/Go-Language-Golang-/controller"
)

func Server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Active).Methods("GET")
	router.HandleFunc("/api/user", controller.GetUser).Methods("GET")
	router.HandleFunc("/api/user", controller.PostUser).Methods("POST")
	router.HandleFunc("/api/user", controller.PatchUser).Methods("PATCH")
	router.HandleFunc("/api/user/singlePhoto", controller.PatchSinglePhoto).Methods("PATCH")
	return router
}
