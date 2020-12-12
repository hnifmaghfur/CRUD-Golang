package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//test api
func Active(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	var status int
	status = 200
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": "Hallo Golang",
	})
}

//get user
func getUser(writer http.ResponseWriter, request *http.Request) {

	var user Users
	var userData []Users //tampung hasil akhir
	var response Response

	rows, err := mysqlDB.Query("SELECT id, name, age FROM users")
	if err != nil {
		renderJson(writer, map[string]interface{}{
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

	response.Status = 200
	response.Message = "Success get Data"
	response.Data = userData

	renderJson(writer, response)
}

//post user
func postUser(writer http.ResponseWriter, request *http.Request) {

	var response Response

	//err := request.ParseMultipartForm(4069)
	//if err != nil {
	//	log.Print(err)
	//}

	parseForm(writer, request, 4069, 500, "Invalid server.")

	name := request.FormValue("name")
	age := request.FormValue("age")

	_, err := mysqlDB.Exec("INSERT INTO users (name, age) values (?,?)", name, age)
	if err != nil {
		renderJson(writer, handleMessage(500, "Error query db"))
		return
	}

	response.Status = 201
	response.Message = "Success add Data"

	renderJson(writer, response)

}

func patchUser(writer http.ResponseWriter, request *http.Request) {
	var response Response

	//get id from params
	getId, ok := request.URL.Query()["id"]
	if !ok {
		response.Status = 404
		response.Message = "Failed to get ID"

		renderJson(writer, response)
		return
	}

	//Parse multipart form
	parseForm(writer, request, 4069, 500, "Invalid server.")

	//store id from params to id
	id := getId[0]
	name := request.FormValue("name")
	age := request.FormValue("age")

	_, err := mysqlDB.Exec("UPDATE users SET name =?, age =? WHERE id=?", name, age, id)
	if err != nil {
		renderJson(writer, handleMessage(500, "Invalid Server."))
		return
	}

	renderJson(writer, handleMessage(200, "Success Patch Data"))
}

func patchSinglePhoto(writer http.ResponseWriter, request *http.Request) {

	//get id from params
	getId, ok := request.URL.Query()["id"]
	if !ok {
		renderJson(writer, handleMessage(500, "internal server error."))
		return
	}

	parseForm(writer, request, 10<<20, 500, "Invalid Server.")

	//get data from frontend
	file, handler, err := request.FormFile("photo")
	if err != nil {
		renderJson(writer, handleMessage(500, "Failed get Data."))
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
		renderJson(writer, handleMessage(500, "failed save photo."))
		return
	}
	//close save photo from folder
	defer storePhoto.Close()

	//save photo from file uploader to images folder
	savePhoto, err := ioutil.ReadAll(file)
	if err != nil {
		renderJson(writer, handleMessage(500, "Failed save photo"))
		return
	}
	storePhoto.Write(savePhoto)

	//store to database
	id := getId[0]
	photo := storePhoto.Name()

	_, err = mysqlDB.Exec("UPDATE users SET photo=? WHERE id=?", photo, id)
	if err != nil {
		renderJson(writer, handleMessage(500, "Invalid server request."))
	}

	renderJson(writer, handleMessage(200, "Success upload photo"))
}
