package service

import (
	"api1/model"
	"api1/repository"
)

func CreateUser(user model.User1) error {
	return repository.CreateUser(user)

}

func GetUserByID(userID int) (model.User1, error) {
	return repository.GetUserByID(userID)
}

func GetAllUser() ([]model.User1, error) {
	return repository.GetAllUser()

}

func DeleteUserByID(userID int) (bool, error) {
	return repository.DeleteUserByID(userID)
}
