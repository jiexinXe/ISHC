package models

type VideoMonitor struct {
	ID      int    `json:"id" db:"id"`
	URL     string `json:"url" db:"url"`
	Status  string `json:"status" db:"status"`
	Address string `json:"address" db:"address"`
	Owner   string `json:"owner" db:"owner"`
}
