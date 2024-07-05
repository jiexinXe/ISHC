package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
)

func CreateEventTask(eventTask *models.EventTask) (int, error) {
	query := `INSERT INTO event_task (event_id, task_id, url) VALUES (?, ?, ?)`
	result, err := config.DB.Exec(query, eventTask.EventID, eventTask.TaskID, eventTask.URL)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func GetEventTaskByID(id int) (*models.EventTask, error) {
	query := `SELECT id, event_id, task_id, url FROM event_task WHERE id = ?`
	row := config.DB.QueryRow(query, id)

	var eventTask models.EventTask
	if err := row.Scan(&eventTask.ID, &eventTask.EventID, &eventTask.TaskID, &eventTask.URL); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &eventTask, nil
}

func UpdateEventTask(eventTask *models.EventTask) error {
	query := `UPDATE event_task SET event_id = ?, task_id = ?, url = ? WHERE id = ?`
	_, err := config.DB.Exec(query, eventTask.EventID, eventTask.TaskID, eventTask.URL, eventTask.ID)
	return err
}

func DeleteEventTask(id int) error {
	query := `DELETE FROM event_task WHERE id = ?`
	_, err := config.DB.Exec(query, id)
	return err
}

func GetAllEventTasks() ([]models.EventTask, error) {
	query := `SELECT id, event_id, task_id, url FROM event_task`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventTasks []models.EventTask
	for rows.Next() {
		var eventTask models.EventTask
		if err := rows.Scan(&eventTask.ID, &eventTask.EventID, &eventTask.TaskID, &eventTask.URL); err != nil {
			return nil, err
		}
		eventTasks = append(eventTasks, eventTask)
	}
	return eventTasks, nil
}
