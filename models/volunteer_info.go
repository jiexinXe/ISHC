package models

type VolunteerInfo struct {
	ID           int        `json:"id" db:"id"`
	OrgID        int        `json:"org_id" db:"org_id"`
	ClientID     int        `json:"client_id" db:"client_id"`
	Name         string     `json:"name" db:"name"`
	Gender       string     `json:"gender" db:"gender"`
	Phone        string     `json:"phone" db:"phone"`
	IDCard       string     `json:"id_card" db:"id_card"`
	Birthday     CustomTime `json:"birthday" db:"birthday"`
	CheckinDate  CustomTime `json:"checkin_date" db:"checkin_date"`
	CheckoutDate CustomTime `json:"checkout_date" db:"checkout_date"`
	ImgsetDir    string     `json:"imgset_dir" db:"imgset_dir"`
	ProfilePhoto string     `json:"profile_photo" db:"profile_photo"`
	Description  string     `json:"description" db:"description"`
	IsActive     string     `json:"is_active" db:"is_active"`
	Created      CustomTime `json:"created" db:"created"`
	CreatedBy    int        `json:"created_by" db:"created_by"`
	Updated      CustomTime `json:"updated" db:"updated"`
	UpdatedBy    int        `json:"updated_by" db:"updated_by"`
	Remove       string     `json:"remove" db:"remove"`
}
