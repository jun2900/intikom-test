package routes

import (
	"intikom-interview/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	taskGroup := router.Group("/tasks")

	userGroup.POST("/", controllers.RegisterUser)
	userGroup.GET("/", controllers.GetAllUsers)
	userGroup.GET("/:id", controllers.GetUserByID)
	userGroup.PUT("/:id", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)

	taskGroup.POST("/", controllers.CreateTask)
	taskGroup.GET("/", controllers.GetAllTasks)
	taskGroup.GET("/:id", controllers.GetTaskByID)
	taskGroup.PUT("/:id", controllers.UpdateTask)
	taskGroup.DELETE("/:id", controllers.DeleteTask)
}
