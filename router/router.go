package router

import (
	"ISHC/controllers"
	"ISHC/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 管理员登录接口
	r.POST("/admin/login", controllers.AdminLogin)

	// 使用JWT中间件保护以下接口
	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())

	// 工作人员接口
	auth.GET("/employees/:id", controllers.GetEmployeeById)
	auth.POST("/employees", controllers.CreateEmployee)
	auth.PUT("/employees/:id", controllers.UpdateEmployee)
	auth.DELETE("/employees/:id", controllers.DeleteEmployee)
	auth.POST("/employees/:id/profile_photo", controllers.SetEmployeeProfilePhoto)

	// 老人接口
	auth.GET("/oldpersons/:id", controllers.GetOldPersonById)
	auth.POST("/oldpersons", controllers.CreateOldPerson)
	auth.PUT("/oldpersons/:id", controllers.UpdateOldPerson)
	auth.DELETE("/oldpersons/:id", controllers.DeleteOldPerson)
	auth.POST("/oldpersons/:id/profile_photo", controllers.SetOldPersonProfilePhoto)

	// 管理员接口
	auth.GET("/admin/:id", controllers.GetAdminById)
	auth.PUT("/admin/:id", controllers.UpdateAdmin)

	// 义工信息接口
	auth.GET("/volunteers/:id", controllers.GetVolunteerById)
	auth.POST("/volunteers", controllers.CreateVolunteer)
	auth.PUT("/volunteers/:id", controllers.UpdateVolunteer)
	auth.DELETE("/volunteers/:id", controllers.DeleteVolunteer)
	auth.POST("/volunteers/:id/profile_photo", controllers.SetVolunteerProfilePhoto)

	// 事件接口
	auth.POST("/events", controllers.CreateEvent)
	auth.GET("/events/type/:type", controllers.GetEventsByType)
	auth.GET("/events/oldperson/:oldperson_id", controllers.GetEventsByOldPersonId)
}
