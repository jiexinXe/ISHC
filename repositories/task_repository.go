package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"fmt"
)

// 添加任务记录
func CreateTask(task *models.Task) error {
	query := `INSERT INTO task (task_type, start_time, end_time, status, camera_id) VALUES (?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, task.TaskType, task.StartTime.Time.Format(models.CtLayoutDateTime), task.EndTime.Time.Format(models.CtLayoutDateTime), task.Status, task.CameraID)
	return err
}

// 查询所有任务记录
func GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, task_type, start_time, end_time, status, camera_id FROM task`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying task: %v", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.TaskType,
			&task.StartTime,
			&task.EndTime,
			&task.Status,
			&task.CameraID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %v", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over task rows: %v", err)
	}

	return tasks, nil
}

// 根据ID更新任务记录
func UpdateTask(task *models.Task) error {
	query := `UPDATE task SET task_type=?, start_time=?, end_time=?, status=?, camera_id=? WHERE id=?`
	_, err := config.DB.Exec(query, task.TaskType, task.StartTime.Time.Format(models.CtLayoutDateTime), task.EndTime.Time.Format(models.CtLayoutDateTime), task.Status, task.CameraID, task.ID)
	return err
}

// 根据ID删除任务记录
func DeleteTask(id int) error {
	query := `DELETE FROM task WHERE id = ?`
	_, err := config.DB.Exec(query, id)
	return err
}
