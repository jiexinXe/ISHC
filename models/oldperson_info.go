package models

type OldPersonInfo struct {
	ID                        int        `json:"id" db:"id"`
	Username                  string     `json:"username" db:"username"`
	Gender                    string     `json:"gender" db:"gender"`
	Phone                     string     `json:"phone" db:"phone"`
	IDCard                    string     `json:"id_card" db:"id_card"`
	Birthday                  CustomTime `json:"birthday" db:"birthday"`
	CheckinDate               CustomTime `json:"checkin_date" db:"checkin_date"`
	CheckoutDate              CustomTime `json:"checkout_date" db:"checkout_date"`
	ImgsetDir                 string     `json:"imgset_dir" db:"imgset_dir"`
	ProfilePhoto              string     `json:"profile_photo" db:"profile_photo"`
	RoomNumber                string     `json:"room_number" db:"room_number"`
	FirstGuardianName         string     `json:"firstguardian_name" db:"firstguardian_name"`
	FirstGuardianRelationship string     `json:"firstguardian_relationship" db:"firstguardian_relationship"`
	FirstGuardianPhone        string     `json:"firstguardian_phone" db:"firstguardian_phone"`
	FirstGuardianWechat       string     `json:"firstguardian_wechat" db:"firstguardian_wechat"`
	HealthState               string     `json:"health_state" db:"health_state"`
	Description               string     `json:"description" db:"description"`
	IsActive                  string     `json:"isactive" db:"isactive"`
	Created                   CustomTime `json:"created" db:"created"`
	CreatedBy                 int        `json:"created_by" db:"createby"`
	Updated                   CustomTime `json:"updated" db:"updated"`
	UpdatedBy                 int        `json:"updated_by" db:"updateby"`
	Remove                    string     `json:"remove" db:"remove"`
}
