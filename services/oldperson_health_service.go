package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateOldPersonHealth(health *models.OldPersonHealth) error {
	return repositories.CreateOldPersonHealth(health)
}

func GetAllOldPersonHealth() ([]models.OldPersonHealth, error) {
	return repositories.GetAllOldPersonHealth()
}

func GetOldPersonHealthByTime(start, end string) ([]models.OldPersonHealth, error) {
	return repositories.GetOldPersonHealthByTime(start, end)
}
