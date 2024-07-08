package controllers

import (
	"ISHC/config"
	"ISHC/models"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

// AdminLogin godoc
// @Summary Login
// @Description Login to the admin account
// @Tags admin
// @Accept json
// @Produce json
// @Param login body models.SysUser true "Login payload"
// @Router /admin/login [post]
func AdminLogin(c *gin.Context) {
	var login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := services.AuthenticateAdmin(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := config.GenerateToken(admin.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"id":       admin.ID,
		"username": admin.UserName,
		"realName": admin.RealName,
		"email":    admin.Email,
	})
}

// GetAdminById godoc
// @Summary Get admin by ID
// @Description Get details of an admin by ID
// @Tags admin
// @Produce json
// @Param id path int true "Admin ID"
// @Router /admin/{id} [get]
func GetAdminById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	admin, err := services.GetAdminById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if admin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// UpdateAdmin godoc
// @Summary Update admin
// @Description Update details of an admin
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Admin ID"
// @Param admin body models.SysUser true "Admin payload"
// @Router /admin/{id} [put]
func UpdateAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var admin models.SysUser
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin.ID = id

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	admin.Password = string(hashedPassword)

	if err := services.UpdateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// GetAdminInfo godoc
// @Summary Get admin info
// @Description Get details of the authenticated admin
// @Tags admin
// @Produce json
// @Router /admin/info [get]
func GetAdminInfo(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	admin, err := services.GetAdminByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if admin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}
