package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateEventTask(eventTask *models.EventTask) (int, error) {
	return repositories.CreateEventTask(eventTask)
}

func GetEventTaskByID(id int) (*models.EventTask, error) {
	return repositories.GetEventTaskByID(id)
}

func UpdateEventTask(eventTask *models.EventTask) error {
	return repositories.UpdateEventTask(eventTask)
}

func DeleteEventTask(id int) error {
	return repositories.DeleteEventTask(id)
}

func GetAllEventTasks() ([]models.EventTask, error) {
	return repositories.GetAllEventTasks()
}
