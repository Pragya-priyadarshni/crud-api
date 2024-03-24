package model

type RequestCreateUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	ID      int    `json:"id"`
}
