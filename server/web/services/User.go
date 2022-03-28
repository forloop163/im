package services

import (
	"gorm.io/gorm"
	"im-project/db"
	"im-project/models"
)

func UserFindByNamePassword(name string, password string) (*gorm.DB, models.User) {
	var user models.User
	result := db.GetDB().Select([]string{"id", "name", "mobile", "email"}).
		Where("name=? AND password=?", name, password).First(&user)
	return result, user
}
