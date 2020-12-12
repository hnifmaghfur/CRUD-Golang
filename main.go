package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

//deklarasi database
var mysqlDB *sql.DB

func main() {
	mysqlDB = connect()
	defer mysqlDB.Close()
	router := server()
	fmt.Printf("Running on port 8010")
	log.Fatal(http.ListenAndServe(":8010", router))
}
