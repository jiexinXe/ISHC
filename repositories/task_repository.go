package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

// 创建任务记录并发送Kafka消息
func CreateTask(task *models.Task) error {
	// 设置默认的 url_string
	if task.URLString == "" {
		task.URLString = fmt.Sprintf("%d/%s", task.CameraID, task.TaskType)
	}

	query := `INSERT INTO task (task_type, start_time, end_time, status, camera_id, url_string) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := config.DB.Exec(query, task.TaskType, task.StartTime.Time.Format(models.CtLayoutDateTime), task.EndTime.Time.Format(models.CtLayoutDateTime), task.Status, task.CameraID, task.URLString)
	if err != nil {
		return fmt.Errorf("无法插入任务记录: %v", err)
	}

	// 获取新插入任务的ID
	taskID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("无法获取新任务的ID: %v", err)
	}
	task.ID = int(taskID)

	// 查询对应的摄像头信息
	camera, err := GetCameraByID(task.CameraID)
	if err != nil {
		return err
	}

	// 构建消息
	taskWithCameraInfo := models.TaskWithCameraInfo{
		Task:   *task,
		Camera: *camera,
	}
	message, err := json.Marshal(taskWithCameraInfo)
	if err != nil {
		return fmt.Errorf("无法序列化消息: %v", err)
	}

	// 设置发送Kafka消息的url
	taskWithCameraInfo.Task.URLString = fmt.Sprintf("rtsp://47.93.76.253:8554/%s", task.URLString)

	// 发送Kafka消息
	err = SendKafkaMessage("video_task_start", message)
	if err != nil {
		return fmt.Errorf("无法发送Kafka消息: %v", err)
	}

	return nil
}

// 根据ID查询摄像头信息
func GetCameraByID(id int) (*models.VideoMonitor, error) {
	query := `SELECT id, url, status, address, owner FROM video_monitor WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var camera models.VideoMonitor
	err := row.Scan(&camera.ID, &camera.URL, &camera.Status, &camera.Address, &camera.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("摄像头ID %d 不存在", id)
		}
		return nil, err
	}
	return &camera, nil
}

// 发送Kafka消息
func SendKafkaMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	_, _, err := config.KafkaProducer.SendMessage(msg)
	return err
}

// 结束任务记录并发送Kafka消息
func FinishTask(taskID int) error {
	// 更新任务状态为 finish
	query := `UPDATE task SET status='finish' WHERE id=?`
	_, err := config.DB.Exec(query, taskID)
	if err != nil {
		return fmt.Errorf("无法更新任务状态: %v", err)
	}

	// 获取任务信息
	task, err := GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("无法获取任务信息: %v", err)
	}

	// 获取对应的摄像头信息
	camera, err := GetCameraByID(task.CameraID)
	if err != nil {
		return fmt.Errorf("无法获取摄像头信息: %v", err)
	}

	// 构建消息
	taskWithCameraInfo := models.TaskWithCameraInfo{
		Task:   *task,
		Camera: *camera,
	}
	message, err := json.Marshal(taskWithCameraInfo)
	if err != nil {
		return fmt.Errorf("无法序列化消息: %v", err)
	}

	// 设置发送Kafka消息的url
	taskWithCameraInfo.Task.URLString = fmt.Sprintf("rtsp://47.93.76.253:8554/%s", task.URLString)

	// 发送Kafka消息
	err = SendKafkaMessage("video_task_end", message)
	if err != nil {
		return fmt.Errorf("无法发送Kafka消息: %v", err)
	}

	return nil
}

// 根据ID查询任务信息
func GetTaskByID(id int) (*models.Task, error) {
	query := `SELECT id, task_type, start_time, end_time, status, camera_id, url_string FROM task WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.TaskType, &task.StartTime, &task.EndTime, &task.Status, &task.CameraID, &task.URLString)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("任务ID %d 不存在", id)
		}
		return nil, err
	}
	return &task, nil
}

// 查询所有任务记录
func GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, task_type, start_time, end_time, status, camera_id, url_string FROM task`
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
			&task.URLString,
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
	query := `UPDATE task SET task_type=?, start_time=?, end_time=?, status=?, camera_id=?, url_string=? WHERE id=?`
	_, err := config.DB.Exec(query, task.TaskType, task.StartTime.Time.Format(models.CtLayoutDateTime), task.EndTime.Time.Format(models.CtLayoutDateTime), task.Status, task.CameraID, task.URLString, task.ID)
	return err
}

// 根据ID删除任务记录
func DeleteTask(id int) error {
	query := `DELETE FROM task WHERE id = ?`
	_, err := config.DB.Exec(query, id)
	return err
}
