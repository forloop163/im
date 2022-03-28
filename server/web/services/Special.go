package services

import (
	"im-project/db"
	"im-project/models"
)

func SpecialGetById(iid uint64) models.Special {
	var special models.Special

	DB := db.GetDB()
	DB.First(&special, iid)

	return special
}

func SpecialWithDetailById(iid uint64) models.Special {
	var special models.Special
	DB := db.GetDB()
	DB.Preload("Detail").Find(&special, iid)

	return special
}
