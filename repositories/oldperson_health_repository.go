package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"fmt"
	"time"
)

func CreateOldPersonHealth(health *models.OldPersonHealth) error {
	// 获取当前时间并设置为 health.Timestamp
	currentTime := time.Now()
	health.Timestamp = models.CustomTime{Time: currentTime}

	query := `INSERT INTO oldperson_health (oldperson_id, heart_rate, timestamp) VALUES (?, ?, ?)`
	_, err := config.DB.Exec(query, health.OldPersonID, health.HeartRate, health.Timestamp.Time.Format(models.CtLayoutDateTime))
	if err != nil {
		return fmt.Errorf("无法插入健康记录: %v", err)
	}
	return nil
}

func GetAllOldPersonHealth() ([]models.OldPersonHealth, error) {
	query := `SELECT id, oldperson_id, heart_rate, timestamp FROM oldperson_health`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying oldperson_health: %v", err)
	}
	defer rows.Close()

	var healthRecords []models.OldPersonHealth
	for rows.Next() {
		var health models.OldPersonHealth
		err := rows.Scan(&health.ID, &health.OldPersonID, &health.HeartRate, &health.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("error scanning oldperson_health row: %v", err)
		}
		healthRecords = append(healthRecords, health)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over oldperson_health rows: %v", err)
	}

	return healthRecords, nil
}

func GetOldPersonHealthByTime(start, end string) ([]models.OldPersonHealth, error) {
	query := `SELECT id, oldperson_id, heart_rate, timestamp FROM oldperson_health WHERE timestamp BETWEEN ? AND ?`
	rows, err := config.DB.Query(query, start, end)
	if err != nil {
		return nil, fmt.Errorf("error querying oldperson_health by time: %v", err)
	}
	defer rows.Close()

	var healthRecords []models.OldPersonHealth
	for rows.Next() {
		var health models.OldPersonHealth
		err := rows.Scan(&health.ID, &health.OldPersonID, &health.HeartRate, &health.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("error scanning oldperson_health row: %v", err)
		}
		healthRecords = append(healthRecords, health)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over oldperson_health rows: %v", err)
	}

	return healthRecords, nil
}
