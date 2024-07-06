package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateEvent(event *models.EventInfo) error {
	return repositories.CreateEvent(event)
}

func GetEventsByType(eventType int) ([]*models.EventInfo, error) {
	return repositories.GetEventsByType(eventType)
}

func GetEventsByOldPersonId(oldPersonId int) ([]*models.EventInfo, error) {
	return repositories.GetEventsByOldPersonId(oldPersonId)
}

func GetAllEvents() ([]models.EventInfo, error) {
	return repositories.GetAllEvents()
}

func GetEventsByTaskId(taskId int) ([]*models.EventInfo, error) {
	return repositories.GetEventsByTaskId(taskId)
}
