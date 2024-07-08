package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateVideoMonitor godoc
// @Summary Create a new video monitor record
// @Description Create a new video monitor record with the input payload
// @Tags videomonitor
// @Accept json
// @Produce json
// @Param video body models.VideoMonitor true "Video Monitor payload"
// @Router /video_monitors [post]
func CreateVideoMonitor(c *gin.Context) {
	var video models.VideoMonitor
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateVideoMonitor(&video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video monitor created successfully"})
}

// GetAllVideoMonitors godoc
// @Summary Get all video monitor records
// @Description Get details of all video monitor records
// @Tags videomonitor
// @Produce json
// @Success 200 {array} models.VideoMonitor
// @Router /video_monitors [get]
func GetAllVideoMonitors(c *gin.Context) {
	videos, err := repositories.GetAllVideoMonitors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

// DeleteVideoMonitor godoc
// @Summary Delete a video monitor record
// @Description Delete a video monitor record by ID
// @Tags videomonitor
// @Produce json
// @Param id path int true "Video Monitor ID"
// @Router /video_monitors/{id} [delete]
func DeleteVideoMonitor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := repositories.DeleteVideoMonitor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video monitor deleted successfully"})
}

// GetVolunteerCount godoc
// @Summary Get volunteer count
// @Description Get the total number of volunteers
// @Tags volunteer
// @Produce json
// @Router /volunteers/count [get]
func GetVolunteerCount(c *gin.Context) {
	count, err := services.GetVolunteerCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
