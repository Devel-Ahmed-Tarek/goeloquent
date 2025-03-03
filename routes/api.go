package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/username/goeloquent/app/controllers" // تأكد من تعديل المسار حسب مشروعك
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users/:id", controllers.GetUserWithPosts)
		api.GET("/users", controllers.GetUsers)
		api.GET("/products", controllers.GetProducts) // مثال على Route للمنتجات (يمكنك إضافته لاحقاً)
	}
}
