package services

import "ISHC/repositories"

func GetTaskStatusCounts() (*repositories.TaskStatusCounts, error) {
	return repositories.GetTaskStatusCounts()
}
