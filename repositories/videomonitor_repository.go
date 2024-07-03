package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"fmt"
)

// 添加视频监控记录
func CreateVideoMonitor(video *models.VideoMonitor) error {
	query := `INSERT INTO video_monitor (url, status, address, owner) VALUES (?, ?, ?, ?)`
	_, err := config.DB.Exec(query, video.URL, video.Status, video.Address, video.Owner)
	return err
}

// 查询所有视频监控记录
func GetAllVideoMonitors() ([]models.VideoMonitor, error) {
	query := `SELECT id, url, status, address, owner FROM video_monitor`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying video_monitor: %v", err)
	}
	defer rows.Close()

	var videos []models.VideoMonitor
	for rows.Next() {
		var video models.VideoMonitor
		err := rows.Scan(
			&video.ID,
			&video.URL,
			&video.Status,
			&video.Address,
			&video.Owner,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning video_monitor row: %v", err)
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over video_monitor rows: %v", err)
	}

	return videos, nil
}

// 根据ID删除视频监控记录
func DeleteVideoMonitor(id int) error {
	query := `DELETE FROM video_monitor WHERE id = ?`
	_, err := config.DB.Exec(query, id)
	return err
}
