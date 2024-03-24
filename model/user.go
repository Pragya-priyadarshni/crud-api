package model

import (
	"gorm.io/gorm"
)

type User1 struct {
	gorm.Model

	Name    string
	Age     int
	Address string
	Id      int
}
