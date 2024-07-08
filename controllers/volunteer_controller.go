package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateVolunteer godoc
// @Summary Create a new volunteer
// @Description Create a new volunteer with the input payload
// @Tags volunteer
// @Accept json
// @Produce json
// @Param volunteer body models.VolunteerInfo true "Volunteer payload"
// @Router /volunteers [post]
func CreateVolunteer(c *gin.Context) {
	var volunteer models.VolunteerInfo
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateVolunteer(&volunteer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, volunteer)
}

// UpdateVolunteer godoc
// @Summary Update a volunteer
// @Description Update a volunteer with the given ID
// @Tags volunteer
// @Accept json
// @Produce json
// @Param id path int true "Volunteer ID"
// @Param volunteer body models.VolunteerInfo true "Volunteer payload"
// @Success 200 {object} models.VolunteerInfo
// @Router /volunteers/{id} [put]
func UpdateVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var volunteer models.VolunteerInfo
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	volunteer.ID = id
	if err := services.UpdateVolunteer(&volunteer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, volunteer)
}

// GetAllVolunteers godoc
// @Summary Get all volunteers
// @Description Get details of all volunteers
// @Tags volunteer
// @Produce json
// @Success 200 {array} models.VolunteerInfo
// @Router /volunteers [get]
func GetAllVolunteers(c *gin.Context) {
	volunteers, err := repositories.GetAllVolunteers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, volunteers)
}

// GetVolunteerById godoc
// @Summary Get a volunteer by ID
// @Description Get details of a volunteer by ID
// @Tags volunteer
// @Produce json
// @Param id path int true "Volunteer ID"
// @Success 200 {object} models.VolunteerInfo
// @Router /volunteers/{id} [get]
func GetVolunteerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	volunteer, err := services.GetVolunteerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if volunteer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
		return
	}
	c.JSON(http.StatusOK, volunteer)
}

// DeleteVolunteer godoc
// @Summary Delete a volunteer
// @Description Delete a volunteer by ID
// @Tags volunteer
// @Produce json
// @Param id path int true "Volunteer ID"
// @Router /volunteers/{id} [delete]
func DeleteVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := services.DeleteVolunteer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Volunteer deleted successfully"})
}

// SetVolunteerProfilePhoto godoc
// @Summary Set volunteer profile photo
// @Description Set profile photo for a volunteer by ID
// @Tags volunteer
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Volunteer ID"
// @Param profile_photo formData string true "Profile Photo URL"
// @Router /volunteers/{id}/profile_photo [post]
func SetVolunteerProfilePhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	profilePhoto := c.PostForm("profile_photo")
	if err := services.SetVolunteerProfilePhoto(id, profilePhoto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile photo updated successfully"})
}
