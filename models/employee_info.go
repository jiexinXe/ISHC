package models

import "time"

type EmployeeInfo struct {
	ID           int       `json:"id" db:"id"`
	OrgID        int       `json:"org_id" db:"org_id"`
	ClientID     int       `json:"client_id" db:"client_id"`
	Username     string    `json:"username" db:"username"`
	Gender       string    `json:"gender" db:"gender"`
	Phone        string    `json:"phone" db:"phone"`
	IDCard       string    `json:"id_card" db:"id_card"`
	Birthday     time.Time `json:"birthday" db:"birthday"`
	HireDate     time.Time `json:"hire_date" db:"hire_date"`
	ResignDate   time.Time `json:"resign_date" db:"resign_date"`
	ImgsetDir    string    `json:"imgset_dir" db:"imgset_dir"`
	ProfilePhoto string    `json:"profile_photo" db:"profile_photo"`
	Description  string    `json:"description" db:"description"`
	IsActive     string    `json:"is_active" db:"is_active"`
	Created      time.Time `json:"created" db:"created"`
	CreatedBy    int       `json:"created_by" db:"created_by"`
	Updated      time.Time `json:"updated" db:"updated"`
	UpdatedBy    int       `json:"updated_by" db:"updated_by"`
	Remove       string    `json:"remove" db:"remove"`
}
