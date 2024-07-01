package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SearchEvents(c *gin.Context) {
	params := make(map[string]string)
	if eventType := c.Query("event_type"); eventType != "" {
		params["event_type"] = eventType
	}
	if eventDate := c.Query("event_date"); eventDate != "" {
		params["event_date"] = eventDate
	}
	if eventLocation := c.Query("event_location"); eventLocation != "" {
		params["event_location"] = eventLocation
	}
	if eventDesc := c.Query("event_desc"); eventDesc != "" {
		params["event_desc"] = eventDesc
	}
	if oldPersonID := c.Query("oldperson_id"); oldPersonID != "" {
		params["oldperson_id"] = oldPersonID
	}

	events, err := repositories.SearchEvents(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

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

func GetAllEvents(c *gin.Context) {
	events, err := repositories.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
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
