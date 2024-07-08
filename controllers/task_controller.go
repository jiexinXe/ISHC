package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task payload"
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

// FinishTask godoc
// @Summary Finish a task
// @Description Mark a task as finished by ID
// @Tags task
// @Produce json
// @Param id path int true "Task ID"
// @Router /tasks/{id}/finish [put]
func FinishTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := repositories.FinishTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task finished successfully"})
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get details of all tasks
// @Tags task
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetAllTasks(c *gin.Context) {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update a task with the input payload
// @Tags task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task payload"
// @Router /tasks [put]
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags task
// @Produce json
// @Param id path int true "Task ID"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := repositories.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// GetTaskStatusCounts godoc
// @Summary Get task status counts
// @Description Get the counts of tasks by their status
// @Tags task
// @Produce json
// @Success 200 {object} map[string]int
// @Router /tasks/status_counts [get]
func GetTaskStatusCounts(c *gin.Context) {
	counts, err := services.GetTaskStatusCounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}
