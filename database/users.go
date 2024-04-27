package database

import (
	"playground/models"
	"playground/request"
	"time"
)

func CreateUser(req request.CreateUserRequest) (models.User, error) {
	var err error
	user := models.User{
		Username:  req.Username,
		Role:      req.Role,
		CreatedAt: time.Now(),
		Password:  req.Password,
	}
	if err = db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, err
}

func GetUser(req request.GetUserRequest) (models.User, error) {
	var err error
	user := models.User{
		Id: req.ID,
	}
	if err = db.First(&user).Error; err != nil {
		return user, err
	}
	return user, err
}

func GetUserLists(req request.GetUserListsRequest) ([]models.User, error) {
	var err error
	user := []models.User{}
	offset := (req.PageNo - 1) * req.PageSize
	if err = db.Limit(req.PageSize).Offset(offset).Find(&user).Error; err != nil {
		return user, err
	}
	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	var err error
	user := models.User{}
	if err = db.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return user, err
	}
	return user, err
}
