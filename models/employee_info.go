package models

type EmployeeInfo struct {
	ID           int        `json:"id" db:"id"`
	OrgID        int        `json:"org_id" db:"org_id"`
	ClientID     int        `json:"client_id" db:"client_id"`
	Username     string     `json:"username" db:"username"`
	Gender       string     `json:"gender" db:"gender"`
	Phone        string     `json:"phone" db:"phone"`
	IDCard       string     `json:"id_card" db:"id_card"`
	Birthday     CustomTime `json:"birthday" db:"birthday"`
	HireDate     CustomTime `json:"hire_date" db:"hire_date"`
	ResignDate   CustomTime `json:"resign_date" db:"resign_date"`
	ImgsetDir    string     `json:"imgset_dir" db:"imgset_dir"`
	ProfilePhoto string     `json:"profile_photo" db:"profile_photo"`
	Description  string     `json:"description" db:"description"`
	IsActive     string     `json:"isactive" db:"isactive"`
	Created      CustomTime `json:"created" db:"created"`
	CreatedBy    int        `json:"created_by" db:"createBY"`
	Updated      CustomTime `json:"updated" db:"updated"`
	UpdatedBy    int        `json:"updated_by" db:"updateby"`
	Remove       string     `json:"remove" db:"remove"`
}
