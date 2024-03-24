package repository

import (
	"api1/db"
	"api1/model"
)

func CreateUser(user model.User1) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllUser() ([]model.User1, error) {
	var listofUsers []model.User1
	result := db.DB.Find(&listofUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	return listofUsers, nil
}

func GetUserByID(userID int) (model.User1, error) {
	var user model.User1
	result := db.DB.First(&user, userID)

	if result.Error != nil {
		return model.User1{}, result.Error
	}
	return user, nil

}

func DeleteUserByID(userID int) (bool, error) {
	var user model.User1
	result := db.DB.First(&user, userID)

	if result.Error != nil {
		return false, result.Error
	}
	db.DB.Delete(&user)
	return true, nil
}
