package controllers

import (
	"ISHC/models"
	"ISHC/repositories"
	"ISHC/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee with the input payload
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body models.EmployeeInfo true "Employee payload"
// @Success 201 {object} models.EmployeeInfo
// @Router /employees [post]
func CreateEmployee(c *gin.Context) {
	var employee models.EmployeeInfo
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an employee with the given ID
// @Tags employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.EmployeeInfo true "Employee payload"
// @Success 200 {object} models.EmployeeInfo
// @Router /employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var employee models.EmployeeInfo
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.ID = id
	if err := services.UpdateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Get details of all employees
// @Tags employee
// @Produce json
// @Success 200 {array} models.EmployeeInfo
// @Router /employees [get]
func GetAllEmployees(c *gin.Context) {
	employees, err := repositories.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeById godoc
// @Summary Get an employee by ID
// @Description Get details of an employee by ID
// @Tags employee
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.EmployeeInfo
// @Router /employees/{id} [get]
func GetEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	employee, err := services.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags employee
// @Produce json
// @Param id path int true "Employee ID"
// @Router /employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := services.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// SetEmployeeProfilePhoto godoc
// @Summary Set employee profile photo
// @Description Set profile photo for an employee by ID
// @Tags employee
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Employee ID"
// @Param profile_photo formData string true "Profile Photo URL"
// @Router /employees/{id}/profile_photo [post]
func SetEmployeeProfilePhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	profilePhoto := c.PostForm("profile_photo")
	if err := services.SetEmployeeProfilePhoto(id, profilePhoto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile photo updated successfully"})
}

// GetEmployeeCount godoc
// @Summary Get employee count
// @Description Get the total number of employees
// @Tags employee
// @Produce json
// @Router /employees/count [get]
func GetEmployeeCount(c *gin.Context) {
	count, err := services.GetEmployeeCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
