package model

type Users struct {
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	Password []uint8 `json:"password"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
