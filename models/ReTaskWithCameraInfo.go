package models

type TaskWithCameraInfo struct {
	Task   Task         `json:"task"`
	Camera VideoMonitor `json:"camera"`
}
