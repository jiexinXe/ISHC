package models

type OldPersonHealth struct {
	ID          int        `json:"id" db:"id"`
	OldPersonID int        `json:"oldperson_id" db:"oldperson_id"`
	HeartRate   int        `json:"heart_rate" db:"heart_rate"`
	Timestamp   CustomTime `json:"timestamp" db:"timestamp"`
}
