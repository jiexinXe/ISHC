package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 创建视频监控记录
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

// 查询所有视频监控记录
func GetAllVideoMonitors(c *gin.Context) {
	videos, err := repositories.GetAllVideoMonitors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

// 删除视频监控记录
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
