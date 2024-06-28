package controllers

import (
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateEvent(c *gin.Context) {
	var event models.EventInfo
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, event)
}

func GetEventsByType(c *gin.Context) {
	eventType, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type"})
		return
	}
	events, err := services.GetEventsByType(eventType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func GetEventsByOldPersonId(c *gin.Context) {
	oldPersonId, err := strconv.Atoi(c.Param("oldperson_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid old person ID"})
		return
	}
	events, err := services.GetEventsByOldPersonId(oldPersonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
