package models

import "time"

type OldPersonInfo struct {
	ID                         int       `json:"id" db:"id"`
	OrgID                      int       `json:"org_id" db:"org_id"`
	ClientID                   int       `json:"client_id" db:"client_id"`
	Username                   string    `json:"username" db:"username"`
	Gender                     string    `json:"gender" db:"gender"`
	Phone                      string    `json:"phone" db:"phone"`
	IDCard                     string    `json:"id_card" db:"id_card"`
	Birthday                   time.Time `json:"birthday" db:"birthday"`
	CheckinDate                time.Time `json:"checkin_date" db:"checkin_date"`
	CheckoutDate               time.Time `json:"checkout_date" db:"checkout_date"`
	ImgsetDir                  string    `json:"imgset_dir" db:"imgset_dir"`
	ProfilePhoto               string    `json:"profile_photo" db:"profile_photo"`
	RoomNumber                 string    `json:"room_number" db:"room_number"`
	FirstGuardianName          string    `json:"firstguardian_name" db:"firstguardian_name"`
	FirstGuardianRelationship  string    `json:"firstguardian_relationship" db:"firstguardian_relationship"`
	FirstGuardianPhone         string    `json:"firstguardian_phone" db:"firstguardian_phone"`
	FirstGuardianWechat        string    `json:"firstguardian_wechat" db:"firstguardian_wechat"`
	SecondGuardianName         string    `json:"secondguardian_name" db:"secondguardian_name"`
	SecondGuardianRelationship string    `json:"secondguardian_relationship" db:"secondguardian_relationship"`
	SecondGuardianPhone        string    `json:"secondguardian_phone" db:"secondguardian_phone"`
	SecondGuardianWechat       string    `json:"secondguardian_wechat" db:"secondguardian_wechat"`
	HealthState                string    `json:"health_state" db:"health_state"`
	Description                string    `json:"description" db:"description"`
	IsActive                   string    `json:"is_active" db:"is_active"`
	Created                    time.Time `json:"created" db:"created"`
	CreatedBy                  int       `json:"created_by" db:"created_by"`
	Updated                    time.Time `json:"updated" db:"updated"`
	UpdatedBy                  int       `json:"updated_by" db:"updated_by"`
	Remove                     string    `json:"remove" db:"remove"`
}
