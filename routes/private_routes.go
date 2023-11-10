package routes

import (
	"intikom-interview/controllers"
	"intikom-interview/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func PrivateRoutes(router *gin.Engine, oauthConfig *oauth2.Config) {
	taskGroup := router.Group("/tasks")

	taskGroup.Use(middleware.VerifyAccessToken(oauthConfig))

	taskGroup.POST("/", controllers.CreateTask)
	taskGroup.GET("/", controllers.GetAllTasks)
	taskGroup.GET("/:id", controllers.GetTaskByID)
	taskGroup.PUT("/:id", controllers.UpdateTask)
	taskGroup.DELETE("/:id", controllers.DeleteTask)
}
