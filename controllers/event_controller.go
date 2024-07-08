package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SearchEvents godoc
// @Summary Search events
// @Description Search events with the given parameters
// @Tags event
// @Produce json
// @Param event_type query string false "Event type"
// @Param event_date query string false "Event date"
// @Param event_location query string false "Event location"
// @Param event_desc query string false "Event description"
// @Param oldperson_id query string false "Old person ID"
// @Param task_id query string false "Task ID"
// @Success 200 {array} models.EventInfo
// @Router /events/search [get]
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
	if taskID := c.Query("task_id"); taskID != "" {
		params["task_id"] = taskID
	}

	events, err := repositories.SearchEvents(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with the input payload
// @Tags event
// @Accept json
// @Produce json
// @Param event body models.EventInfo true "Event payload"
// @Success 201 {object} models.EventInfo
// @Router /events [post]
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

// GetAllEvents godoc
// @Summary Get all events
// @Description Get details of all events
// @Tags event
// @Produce json
// @Success 200 {array} models.EventInfo
// @Router /events [get]
func GetAllEvents(c *gin.Context) {
	events, err := services.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// GetEventsByType godoc
// @Summary Get events by type
// @Description Get details of events by type
// @Tags event
// @Produce json
// @Param type path int true "Event type"
// @Success 200 {array} models.EventInfo
// @Router /events/type/{type} [get]
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

// GetEventsByOldPersonId godoc
// @Summary Get events by old person ID
// @Description Get details of events by old person ID
// @Tags event
// @Produce json
// @Param oldperson_id path int true "Old person ID"
// @Success 200 {array} models.EventInfo
// @Router /events/oldperson/{oldperson_id} [get]
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

// GetEventsByTaskId godoc
// @Summary Get events by task ID
// @Description Get details of events by task ID
// @Tags event
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 200 {array} models.EventInfo
// @Router /events/task/{task_id} [get]
func GetEventsByTaskId(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	events, err := services.GetEventsByTaskId(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
