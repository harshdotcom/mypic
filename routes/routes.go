package routes

import (
	"mypic/controllers"
	"mypic/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		users := protected.Group("/users")
		{
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
}
