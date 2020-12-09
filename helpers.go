package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)


//connect DB
func connect() *sql.DB {
	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "golang"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}


//formResponse
func renderJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	//log.Print(data)
}


//handle message
func handleMessage(status int, message string) Response {
	var response Response

	response.Status = status
	response.Message = message

	return response
}