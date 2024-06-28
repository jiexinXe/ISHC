package models

type EventInfo struct {
	ID            int        `json:"id" db:"id"`
	EventType     int        `json:"event_type" db:"event_type"`
	EventDate     CustomTime `json:"event_date" db:"event_date"`
	EventLocation string     `json:"event_location" db:"event_location"`
	EventDesc     string     `json:"event_desc" db:"event_desc"`
	OldPersonID   int        `json:"oldperson_id" db:"oldperson_id"`
}
