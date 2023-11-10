package routes

import (
	"intikom-interview/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	//taskGroup := router.Group("/tasks")

	userGroup.POST("/", controllers.RegisterUser)
	userGroup.GET("/", controllers.GetAllUsers)
	userGroup.GET("/:id", controllers.GetUserByID)
	userGroup.PUT("/:id", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)
}
