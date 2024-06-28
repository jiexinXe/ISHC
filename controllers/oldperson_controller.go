package controllers

import (
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOldPerson(c *gin.Context) {
	var oldPerson models.OldPersonInfo
	if err := c.ShouldBindJSON(&oldPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateOldPerson(&oldPerson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, oldPerson)
}

func UpdateOldPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var oldPerson models.OldPersonInfo
	if err := c.ShouldBindJSON(&oldPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oldPerson.ID = id
	if err := services.UpdateOldPerson(&oldPerson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, oldPerson)
}

func GetOldPersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	oldPerson, err := services.GetOldPersonById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if oldPerson == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Old person not found"})
		return
	}
	c.JSON(http.StatusOK, oldPerson)
}

func DeleteOldPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := services.DeleteOldPerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Old person deleted successfully"})
}

func SetOldPersonProfilePhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	profilePhoto := c.PostForm("profile_photo")
	if err := services.SetOldPersonProfilePhoto(id, profilePhoto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile photo updated successfully"})
}
