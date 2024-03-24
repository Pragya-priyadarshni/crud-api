package main

import (
	"api1/router"

	"api1/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectToPostgresDB()
	a := gin.Default()
	router.UserRouter(a)
	a.Run()

}
