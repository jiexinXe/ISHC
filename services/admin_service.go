package services

import (
	"ISHC/models"
	"ISHC/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GetAdminByUsername(username string) (*models.SysUser, error) {
	return repositories.GetAdminByUsername(username)
}

func GetAdminById(id int) (*models.SysUser, error) {
	return repositories.GetAdminById(id)
}

func UpdateAdmin(admin *models.SysUser) error {
	return repositories.UpdateAdmin(admin)
}

func AuthenticateAdmin(username, password string) (*models.SysUser, error) {
	admin, err := repositories.GetAdminByUsername(username)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("admin not found")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return admin, nil
}
