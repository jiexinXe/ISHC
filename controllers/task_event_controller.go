package controllers

import (
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateEventTask godoc
// @Summary Create a new event task
// @Description Create a new event task with the input payload
// @Tags eventtask
// @Accept json
// @Produce json
// @Param eventTask body models.EventTask true "Event Task payload"
// @Success 200 {object} models.EventTask
// @Router /event_tasks [post]
func CreateEventTask(c *gin.Context) {
	var eventTask models.EventTask
	if err := c.ShouldBindJSON(&eventTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateEventTask(&eventTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	eventTask.ID = id

	c.JSON(http.StatusOK, eventTask)
}

// GetEventTaskByID godoc
// @Summary Get an event task by ID
// @Description Get details of an event task by ID
// @Tags eventtask
// @Produce json
// @Param id path int true "Event Task ID"
// @Success 200 {object} models.EventTask
// @Router /event_tasks/{id} [get]
func GetEventTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	eventTask, err := services.GetEventTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if eventTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EventTask not found"})
		return
	}

	c.JSON(http.StatusOK, eventTask)
}

// UpdateEventTask godoc
// @Summary Update an event task
// @Description Update an event task with the input payload
// @Tags eventtask
// @Accept json
// @Produce json
// @Param id path int true "Event Task ID"
// @Param eventTask body models.EventTask true "Event Task payload"
// @Success 200 {object} models.EventTask
// @Router /event_tasks/{id} [put]
func UpdateEventTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var eventTask models.EventTask
	if err := c.ShouldBindJSON(&eventTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	eventTask.ID = id

	if err := services.UpdateEventTask(&eventTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eventTask)
}

// DeleteEventTask godoc
// @Summary Delete an event task
// @Description Delete an event task by ID
// @Tags eventtask
// @Produce json
// @Param id path int true "Event Task ID"
// @Router /event_tasks/{id} [delete]
func DeleteEventTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.DeleteEventTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "EventTask deleted"})
}

// GetAllEventTasks godoc
// @Summary Get all event tasks
// @Description Get details of all event tasks
// @Tags eventtask
// @Produce json
// @Success 200 {array} models.EventTask
// @Router /event_tasks [get]
func GetAllEventTasks(c *gin.Context) {
	eventTasks, err := services.GetAllEventTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eventTasks)
}
