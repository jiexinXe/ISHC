package controllers

import (
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOldPersonHealth godoc
// @Summary Create a new old person health record
// @Description Create a new old person health record with the input payload
// @Tags oldpersonhealth
// @Accept json
// @Produce json
// @Param health body models.OldPersonHealth true "Old Person Health payload"
// @Success 201 {object} models.OldPersonHealth
// @Router /oldperson_health [post]
func CreateOldPersonHealth(c *gin.Context) {
	var health models.OldPersonHealth
	if err := c.ShouldBindJSON(&health); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateOldPersonHealth(&health); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, health)
}

// GetAllOldPersonHealth godoc
// @Summary Get all old person health records
// @Description Get details of all old person health records
// @Tags oldpersonhealth
// @Produce json
// @Success 200 {array} models.OldPersonHealth
// @Router /oldperson_health [get]
func GetAllOldPersonHealth(c *gin.Context) {
	healthRecords, err := services.GetAllOldPersonHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, healthRecords)
}

// GetOldPersonHealthByTime godoc
// @Summary Get old person health records by time range
// @Description Get details of old person health records within the specified time range
// @Tags oldpersonhealth
// @Produce json
// @Param start query string true "Start time"
// @Param end query string true "End time"
// @Success 200 {array} models.OldPersonHealth
// @Router /oldperson_health/time [get]
func GetOldPersonHealthByTime(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start and end time are required"})
		return
	}

	healthRecords, err := services.GetOldPersonHealthByTime(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, healthRecords)
}
