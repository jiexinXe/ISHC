package services

import (
	"ISHC/models"
	"ISHC/repositories"
)

func CreateEmployee(employee *models.EmployeeInfo) error {
	return repositories.CreateEmployee(employee)
}

func UpdateEmployee(employee *models.EmployeeInfo) error {
	return repositories.UpdateEmployee(employee)
}

func GetEmployeeById(id int) (*models.EmployeeInfo, error) {
	return repositories.GetEmployeeById(id)
}

func DeleteEmployee(id int) error {
	return repositories.DeleteEmployee(id)
}

func SetEmployeeProfilePhoto(id int, profilePhoto string) error {
	return repositories.SetEmployeeProfilePhoto(id, profilePhoto)
}

func GetEmployeeCount() (int, error) {
	return repositories.GetEmployeeCount()
}
