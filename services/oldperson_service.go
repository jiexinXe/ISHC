package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateOldPerson(oldPerson *models.OldPersonInfo) error {
	return repositories.CreateOldPerson(oldPerson)
}

func UpdateOldPerson(oldPerson *models.OldPersonInfo) error {
	return repositories.UpdateOldPerson(oldPerson)
}

func GetOldPersonById(id int) (*models.OldPersonInfo, error) {
	return repositories.GetOldPersonById(id)
}

func DeleteOldPerson(id int) error {
	return repositories.DeleteOldPerson(id)
}

func SetOldPersonProfilePhoto(id int, profilePhoto string) error {
	return repositories.SetOldPersonProfilePhoto(id, profilePhoto)
}
