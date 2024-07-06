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
	auth.GET("/employees", controllers.GetAllEmployees)
	auth.GET("/employees/:id", controllers.GetEmployeeById)
	auth.POST("/employees", controllers.CreateEmployee)
	auth.PUT("/employees/:id", controllers.UpdateEmployee)
	auth.DELETE("/employees/:id", controllers.DeleteEmployee)
	auth.POST("/employees/:id/profile_photo", controllers.SetEmployeeProfilePhoto)

	// 老人接口
	auth.GET("/oldpersons", controllers.GetAllOldPersons)
	auth.GET("/oldpersons/:id", controllers.GetOldPersonById)
	auth.POST("/oldpersons", controllers.CreateOldPerson)
	auth.PUT("/oldpersons/:id", controllers.UpdateOldPerson)
	auth.DELETE("/oldpersons/:id", controllers.DeleteOldPerson)
	auth.POST("/oldpersons/:id/profile_photo", controllers.SetOldPersonProfilePhoto)

	// 管理员接口
	auth.GET("/admin/:id", controllers.GetAdminById)
	auth.PUT("/admin/:id", controllers.UpdateAdmin)
	auth.GET("/admin/info", controllers.GetAdminInfo)

	// 义工信息接口
	auth.GET("/volunteers", controllers.GetAllVolunteers)
	auth.GET("/volunteers/:id", controllers.GetVolunteerById)
	auth.POST("/volunteers", controllers.CreateVolunteer)
	auth.PUT("/volunteers/:id", controllers.UpdateVolunteer)
	auth.DELETE("/volunteers/:id", controllers.DeleteVolunteer)
	auth.POST("/volunteers/:id/profile_photo", controllers.SetVolunteerProfilePhoto)

	// 事件接口
	auth.POST("/events", controllers.CreateEvent)
	auth.GET("/events/search", controllers.SearchEvents)
	auth.GET("/events", controllers.GetAllEvents)
	auth.GET("/events/type/:type", controllers.GetEventsByType)
	auth.GET("/events/oldperson/:oldperson_id", controllers.GetEventsByOldPersonId)

	// 监控接口
	auth.POST("/video_monitors", controllers.CreateVideoMonitor)
	auth.GET("/video_monitors", controllers.GetAllVideoMonitors)
	auth.DELETE("/video_monitors/:id", controllers.DeleteVideoMonitor)

	// 任务接口
	auth.POST("/tasks", controllers.CreateTask)
	auth.GET("/tasks", controllers.GetAllTasks)
	auth.PUT("/tasks/:id", controllers.UpdateTask)
	auth.DELETE("/tasks/:id", controllers.DeleteTask)
	auth.PUT("/tasks/:id/finish", controllers.FinishTask)

	// 任务事件接口
	auth.POST("/event_tasks", controllers.CreateEventTask)
	auth.GET("/event_tasks/:id", controllers.GetEventTaskByID)
	auth.PUT("/event_tasks/:id", controllers.UpdateEventTask)
	auth.DELETE("/event_tasks/:id", controllers.DeleteEventTask)
	auth.GET("/event_tasks", controllers.GetAllEventTasks)

}
