package repositories

import (
	"ISHC/config"
	"ISHC/models"
)

func CreateEvent(event *models.EventInfo) error {
	query := `INSERT INTO event_info (event_type, event_date, event_location, event_desc, oldperson_id) 
              VALUES (?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, event.EventType, event.EventDate, event.EventLocation, event.EventDesc, event.OldPersonID)
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
