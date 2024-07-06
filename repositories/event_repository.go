package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"fmt"
	"strings"
)

func SearchEvents(params map[string]string) ([]models.EventInfo, error) {
	baseQuery := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id, image, task_id FROM event_info WHERE`
	var conditions []string
	var args []interface{}

	if val, ok := params["event_type"]; ok {
		conditions = append(conditions, "event_type = ?")
		args = append(args, val)
	}
	if val, ok := params["event_date"]; ok {
		conditions = append(conditions, "event_date = ?")
		args = append(args, val)
	}
	if val, ok := params["event_location"]; ok {
		conditions = append(conditions, "event_location LIKE ?")
		args = append(args, "%"+val+"%")
	}
	if val, ok := params["event_desc"]; ok {
		conditions = append(conditions, "event_desc LIKE ?")
		args = append(args, "%"+val+"%")
	}
	if val, ok := params["oldperson_id"]; ok {
		conditions = append(conditions, "oldperson_id = ?")
		args = append(args, val)
	}
	if val, ok := params["task_id"]; ok {
		conditions = append(conditions, "task_id = ?")
		args = append(args, val)
	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no valid query parameters provided")
	}

	query := baseQuery + " " + strings.Join(conditions, " AND ")

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying event_info: %v", err)
	}
	defer rows.Close()

	var events []models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		err := rows.Scan(
			&event.ID,
			&event.EventType,
			&event.EventDate,
			&event.EventLocation,
			&event.EventDesc,
			&event.OldPersonID,
			&event.Image,
			&event.TaskID)
		if err != nil {
			return nil, fmt.Errorf("error scanning event_info row: %v", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over event_info rows: %v", err)
	}

	return events, nil
}

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

	query := `INSERT INTO event_info (event_type, event_date, event_location, event_desc, oldperson_id, image, task_id) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err = config.DB.Exec(query,
		event.EventType,
		event.EventDate.Time.Format(models.CtLayoutDateTime),
		event.EventLocation,
		event.EventDesc,
		event.OldPersonID,
		event.Image,
		event.TaskID)
	return err
}

func GetEventsByType(eventType int) ([]*models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id, image, task_id 
              FROM event_info WHERE event_type=?`
	rows, err := config.DB.Query(query, eventType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		if err := rows.Scan(&event.ID, &event.EventType, &event.EventDate, &event.EventLocation, &event.EventDesc, &event.OldPersonID, &event.Image, &event.TaskID); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func GetAllEvents() ([]models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id, image, task_id FROM event_info`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying event_info: %v", err)
	}
	defer rows.Close()

	var events []models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		err := rows.Scan(
			&event.ID,
			&event.EventType,
			&event.EventDate,
			&event.EventLocation,
			&event.EventDesc,
			&event.OldPersonID,
			&event.Image,
			&event.TaskID)
		if err != nil {
			return nil, fmt.Errorf("error scanning event_info row: %v", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over event_info rows: %v", err)
	}

	return events, nil
}

func GetEventsByOldPersonId(oldPersonId int) ([]*models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id, image, task_id 
              FROM event_info WHERE oldperson_id=?`
	rows, err := config.DB.Query(query, oldPersonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		if err := rows.Scan(&event.ID, &event.EventType, &event.EventDate, &event.EventLocation, &event.EventDesc, &event.OldPersonID, &event.Image, &event.TaskID); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func GetEventsByTaskId(taskId int) ([]*models.EventInfo, error) {
	query := `SELECT id, event_type, event_date, event_location, event_desc, oldperson_id, image, task_id 
              FROM event_info WHERE task_id=?`
	rows, err := config.DB.Query(query, taskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.EventInfo
	for rows.Next() {
		var event models.EventInfo
		if err := rows.Scan(&event.ID, &event.EventType, &event.EventDate, &event.EventLocation, &event.EventDesc, &event.OldPersonID, &event.Image, &event.TaskID); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}
