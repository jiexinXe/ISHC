package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateVolunteer(volunteer *models.VolunteerInfo) error {
	return repositories.CreateVolunteer(volunteer)
}

func UpdateVolunteer(volunteer *models.VolunteerInfo) error {
	return repositories.UpdateVolunteer(volunteer)
}

func GetVolunteerById(id int) (*models.VolunteerInfo, error) {
	return repositories.GetVolunteerById(id)
}

func DeleteVolunteer(id int) error {
	return repositories.DeleteVolunteer(id)
}

func SetVolunteerProfilePhoto(id int, profilePhoto string) error {
	return repositories.SetVolunteerProfilePhoto(id, profilePhoto)
}

func GetVolunteerCount() (int, error) {
	return repositories.GetVolunteerCount()
}
