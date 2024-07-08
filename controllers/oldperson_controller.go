package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOldPerson godoc
// @Summary Create a new old person
// @Description Create a new old person with the input payload
// @Tags oldperson
// @Accept json
// @Produce json
// @Param oldPerson body models.OldPersonInfo true "Old Person payload"
// @Success 201 {object} models.OldPersonInfo
// @Router /oldpersons [post]
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

// UpdateOldPerson godoc
// @Summary Update an old person
// @Description Update an old person with the given ID
// @Tags oldperson
// @Accept json
// @Produce json
// @Param id path int true "Old Person ID"
// @Param oldPerson body models.OldPersonInfo true "Old Person payload"
// @Success 200 {object} models.OldPersonInfo
// @Router /oldpersons/{id} [put]
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

// GetAllOldPersons godoc
// @Summary Get all old persons
// @Description Get details of all old persons
// @Tags oldperson
// @Produce json
// @Success 200 {array} models.OldPersonInfo
// @Router /oldpersons [get]
func GetAllOldPersons(c *gin.Context) {
	oldPersons, err := repositories.GetAllOldPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, oldPersons)
}

// GetOldPersonById godoc
// @Summary Get an old person by ID
// @Description Get details of an old person by ID
// @Tags oldperson
// @Produce json
// @Param id path int true "Old Person ID"
// @Success 200 {object} models.OldPersonInfo
// @Router /oldpersons/{id} [get]
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

// DeleteOldPerson godoc
// @Summary Delete an old person
// @Description Delete an old person by ID
// @Tags oldperson
// @Produce json
// @Param id path int true "Old Person ID"
// @Router /oldpersons/{id} [delete]
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

// SetOldPersonProfilePhoto godoc
// @Summary Set old person profile photo
// @Description Set profile photo for an old person by ID
// @Tags oldperson
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Old Person ID"
// @Param profile_photo formData string true "Profile Photo URL"
// @Router /oldpersons/{id}/profile_photo [post]
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

// GetOldPersonCount godoc
// @Summary Get old person count
// @Description Get the total number of old persons
// @Tags oldperson
// @Produce json
// @Router /oldpersons/count [get]
func GetOldPersonCount(c *gin.Context) {
	count, err := services.GetOldPersonCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
