package services

import (
	"im-project/db"
	"im-project/models"
)

func UserFindByNamePassword(name string, password string) models.User {
	var user models.User
	db.GetDB().Where("name=? AND password=?", name, password).First(&user)
	return user
}
