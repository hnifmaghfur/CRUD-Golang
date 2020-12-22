package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hnifmaghfur/Go-Language-Golang-/model"
)

//connect DB
func Connect() *sql.DB {
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
func RenderJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

//handle message
func HandleMessage(status int, message string) model.Response {
	var response model.Response

	response.Status = status
	response.Message = message

	return response
}

//parse Form multipart data.
func ParseForm(w http.ResponseWriter, r *http.Request, memory int64, status int, message string) {
	err := r.ParseMultipartForm(memory)
	if err != nil {
		RenderJson(w, HandleMessage(status, message))
	}
}
