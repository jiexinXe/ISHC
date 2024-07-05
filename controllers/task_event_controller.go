package controllers

import (
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func GetAllEventTasks(c *gin.Context) {
	eventTasks, err := services.GetAllEventTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eventTasks)
}
