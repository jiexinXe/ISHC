package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"fmt"
)

func CheckOldPersonExists(oldPersonID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM oldperson_info WHERE id = ?)"
	err := config.DB.QueryRow(query, oldPersonID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("error checking if old person exists: %v", err)
	}
	return exists, nil
}

func CreateEvent(event *models.EventInfo) error {
	// 检查 oldperson_id 是否存在
	exists, err := CheckOldPersonExists(event.OldPersonID)
	if err != nil {
		return fmt.Errorf("error checking if old person exists: %v", err)
	}
	if !exists {
		return fmt.Errorf("old person with id %d does not exist", event.OldPersonID)
	}

	query := `INSERT INTO event_info (event_type, event_date, event_location, event_desc, oldperson_id) 
              VALUES (?, ?, ?, ?, ?)`

	_, err = config.DB.Exec(query,
		event.EventType,
		event.EventDate.Time.Format(models.CtLayoutDateTime),
		event.EventLocation,
		event.EventDesc,
		event.OldPersonID)
	return err
}

func GetEventsByType(eventType int) ([]*models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id 
              FROM event_info WHERE event_type=?`
	rows, err := config.DB.Query(query, eventType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		if err := rows.Scan(&event.ID, &event.EventType, &event.EventDate, &event.EventLocation, &event.EventDesc, &event.OldPersonID); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func GetEventsByOldPersonId(oldPersonId int) ([]*models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id 
              FROM event_info WHERE oldperson_id=?`
	rows, err := config.DB.Query(query, oldPersonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		if err := rows.Scan(&event.ID, &event.EventType, &event.EventDate, &event.EventLocation, &event.EventDesc, &event.OldPersonID); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}
