package model

type Users struct {
	ID   int	`json:"id"`
	Name string	`json:"name"`
	Age  int	`json:"age"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

