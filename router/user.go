package router

import (
	"api1/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	c.GET("/a", handler.GetAllUser)
	c.GET("/b/:id", handler.GetUserByID)
	c.POST("/c", handler.CreateUser)
	c.DELETE("/d/:id", handler.DeleteUserByID)
}
