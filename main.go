package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hnifmaghfur/Go-Language-Golang-/router"
)

func main() {
	router := router.Server()
	fmt.Printf("Running on port 8010")
	log.Fatal(http.ListenAndServe(":8010", router))
}