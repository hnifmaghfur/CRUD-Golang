package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hnifmaghfur/Go-Language-Golang-/helper"
	"github.com/hnifmaghfur/Go-Language-Golang-/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	var user model.Users
	var userData []model.Users
	var tokenAuth interface{}

	//parse mutlipart form
	helper.ParseForm(writer, request, 1026, http.StatusBadGateway, "Bad Getway")

	//get data email && password from form
	email := request.FormValue("email")
	password := request.FormValue("password")

	//open db
	mysqlDB = helper.Connect()
	defer mysqlDB.Close()

	//get data password and id
	rawData, err := mysqlDB.Query("SELECT id, name, password FROM users WHERE email=?", email)
	if err != nil {
		helper.LogView(err, http.StatusBadRequest, "Error")
		helper.RenderJson(writer, helper.HandleMessage(http.StatusBadRequest, "Bad Request"))
		return
	}

	for rawData.Next() {
		if err = rawData.Scan(&user.ID, &user.Name, &user.Password); err != nil {
			helper.LogView(err, http.StatusNotFound, "error")
			helper.RenderJson(writer, helper.HandleMessage(http.StatusBadRequest, "Failed to get Data."))
		} else {
			userData = append(userData, user)
		}
	}

	//convert to byte input password
	hashPassword := []byte(password)

	//compare with bcrypt
	err = bcrypt.CompareHashAndPassword(userData[0].Password, hashPassword)

	if err != nil {
		helper.LogView(err, 404, "error")
		return
	}

	//make secret key in byte for JWT
	secretKey := []byte("superSemarIndonesia")

	//create token JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userData[0].ID,
		"name": userData[0].Name,
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		helper.LogView(err, 404, "Error")
		return
	}

	//use token for Authorization with Bearer
	tokenAuth = "Bearer " + token

	//output token
	helper.LogView(token, 200, "Success")
	helper.RenderJson(writer, tokenAuth)

}
