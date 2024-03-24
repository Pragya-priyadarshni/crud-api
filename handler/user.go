package handler

import (
	"api1/model"
	"api1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var RequestCreateUser model.RequestCreateUser

	err := c.ShouldBindJSON(&RequestCreateUser)

	if err != nil {
		badResponse := model.ReplyCreateUser{
			Message: "Please try again!",
		}
		c.JSON(http.StatusBadRequest, badResponse)
		return
	}
	user := model.User1{
		Name:    RequestCreateUser.Name,
		Age:     RequestCreateUser.Age,
		Id:      RequestCreateUser.ID,
		Address: RequestCreateUser.Address,
	}
	err = service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ReplyCreateUser{
			Message: "Unable to store in DB",
		})

		return
	}
	c.JSON(http.StatusAccepted, model.ReplyCreateUser{
		Message: "User Got Saved!",
	})
}

func GetAllUser(c *gin.Context) {
	listofUsers, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ReplyCreateUser{
			Message: "Unable to store in db",
		})
		return
	}
	c.JSON(http.StatusOK, listofUsers)

}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		badResponse := model.ReplyCreateUser{
			Message: "Unable to store in db",
		}
		c.JSON(http.StatusBadRequest, badResponse)
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ReplyCreateUser{
			Message: "Unable to store in db",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUserByID(c *gin.Context) {
	id := c.Params.ByName("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		badResponse := model.ReplyCreateUser{
			Message: "Unable to store in db",
		}
		c.JSON(http.StatusBadRequest, badResponse)
		return
	}

	ok, err := service.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ReplyCreateUser{
			Message: "Unable to store in db",
		})
		return
	}
	c.JSON(http.StatusOK, ok)
}
