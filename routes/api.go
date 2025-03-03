package routes

import (
	"github.com/Devel-Ahmed-Tarek/goeloquent/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users/:id", controllers.GetUserWithPosts)
		api.GET("/users", controllers.GetUsers)
		// يمكنك إضافة المزيد من الـ routes هنا، مثل /products
	}
}
