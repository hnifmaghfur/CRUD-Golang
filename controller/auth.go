package controller

import (
	"net/http"

	"github.com/hnifmaghfur/Go-Language-Golang-/helper"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	helper.ParseForm(writer, request, 1026, http.StatusBadGateway, "Bad Getway")

	
}