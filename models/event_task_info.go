package models

type EventTask struct {
	ID      int    `json:"id" db:"id"`
	EventID int    `json:"event_id" db:"event_id"`
	TaskID  int    `json:"task_id" db:"task_id"`
	URL     string `json:"url" db:"url"`
}
