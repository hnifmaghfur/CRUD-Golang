package main

import (
	"database/sql"
	"log"
	"net/http"
)

//deklarasi database
var mysqlDB *sql.DB

func main() {
	mysqlDB = connect()
	defer mysqlDB.Close()
	router := server()
	log.Fatal(http.ListenAndServe(":8010", router))
}