package main

type Users struct {
	ID   int
	Name string
	Age  int
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

