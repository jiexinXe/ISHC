package models

type Task struct {
	ID        int        `json:"id" db:"id"`
	TaskType  string     `json:"task_type" db:"task_type"`
	StartTime CustomTime `json:"start_time" db:"start_time"`
	EndTime   CustomTime `json:"end_time" db:"end_time"`
	Status    string     `json:"status" db:"status"`
	CameraID  int        `json:"camera_id" db:"camera_id"`
	URLString string     `json:"url_string" db:"url_string"`
}
