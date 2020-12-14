package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"database/sql"
	"github.com/hnifmaghfur/Go-Language-Golang-/helper"
	"github.com/hnifmaghfur/Go-Language-Golang-/model"
)

//deklarasi database
var mysqlDB *sql.DB

//test api
func Active(w http.ResponseWriter, r *http.Request) {
	status := 200
	w.WriteHeader(status)
	helper.RenderJson(w, helper.HandleMessage(status, "Hallo Golang"))
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"status":  status,
	// 	"message": "Hallo Golang",
	// })
}

//get user
func GetUser(writer http.ResponseWriter, request *http.Request) {

	var user model.Users
	var userData []model.Users //tampung hasil akhir

	mysqlDB = helper.Connect()
	defer mysqlDB.Close()

	rows, err :=  mysqlDB.Query("SELECT id, name, age FROM users")
	if err != nil {
		helper.RenderJson(writer, map[string]interface{}{
			"message": "Not founds.",
		})
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Print(err)
		} else {
			userData = append(userData, user)
		}
	}

	helper.RenderJson(writer, userData)
}

//post user
func PostUser(writer http.ResponseWriter, request *http.Request) {

	var response model.Response

	//err := request.ParseMultipartForm(4069)
	//if err != nil {
	//	log.Print(err)
	//}

	helper.ParseForm(writer, request, 4069, 500, "Invalid server.")

	name := request.FormValue("name")
	age := request.FormValue("age")

	mysqlDB = helper.Connect()
	defer mysqlDB.Close()

	_, err := mysqlDB.Exec("INSERT INTO users (name, age) values (?,?)", name, age)
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "Error query db"))
		return
	}

	response.Status = 201
	response.Message = "Success add Data"

	helper.RenderJson(writer, response)

}

func PatchUser(writer http.ResponseWriter, request *http.Request) {
	var response model.Response

	//get id from params
	getId, ok := request.URL.Query()["id"]
	if !ok {
		response.Status = 404
		response.Message = "Failed to get ID"

		helper.RenderJson(writer, response)
		return
	}

	//Parse multipart form
	helper.ParseForm(writer, request, 4069, 500, "Invalid server.")

	//store id from params to id
	id := getId[0]
	name := request.FormValue("name")
	age := request.FormValue("age")

	mysqlDB = helper.Connect()
	defer mysqlDB.Close()

	_, err := mysqlDB.Exec("UPDATE users SET name =?, age =? WHERE id=?", name, age, id)
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "Invalid Server."))
		return
	}

	helper.RenderJson(writer, helper.HandleMessage(200, "Success Patch Data"))
}

func PatchSinglePhoto(writer http.ResponseWriter, request *http.Request) {

	//get id from params
	getId, ok := request.URL.Query()["id"]
	if !ok {
		helper.RenderJson(writer, helper.HandleMessage(500, "internal server error."))
		return
	}

	helper.ParseForm(writer, request, 10<<20, 500, "Invalid Server.")

	//get data from frontend
	file, handler, err := request.FormFile("photo")
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "Failed get Data."))
		return
	}

	//close get data
	defer file.Close()

	fmt.Printf("Upload file name : %+v\n", handler.Filename)
	fmt.Printf("Size : %+v\n", handler.Size)
	fmt.Printf("MIME Type : %+v\n", handler.Header)

	//store image to images folder
	storePhoto, err := ioutil.TempFile("images", "upload-*.png")
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "failed save photo."))
		return
	}
	//close save photo from folder
	defer storePhoto.Close()

	//save photo from file uploader to images folder
	savePhoto, err := ioutil.ReadAll(file)
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "Failed save photo"))
		return
	}
	_, err = storePhoto.Write(savePhoto)
	if err != nil {
		helper.ParseForm(writer, request, 1024, http.StatusBadRequest, "Failed store photo.")
	}

	//store to database
	id := getId[0]
	photo := storePhoto.Name()

	mysqlDB = helper.Connect()
	defer mysqlDB.Close()

	_, err = mysqlDB.Exec("UPDATE users SET photo=? WHERE id=?", photo, id)
	if err != nil {
		helper.RenderJson(writer, helper.HandleMessage(500, "Invalid server request."))
	}

	helper.RenderJson(writer, helper.HandleMessage(200, "Success upload photo"))
}
