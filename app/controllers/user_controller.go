package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/goeloquent/app/models"  // تأكد من تعديل المسار
	"github.com/username/goeloquent/goeloquent"     // للوصول إلى دوال المكتبة مثل Paginate, ScopeActive, إلخ.
)

// GetUsers يقوم بإرجاع جميع المستخدمين مع دعم Pagination
func GetUsers(c *gin.Context) {
	var users []models.User
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	result, err := goeloquent.Paginate(goeloquent.DB, &models.User{}, &users, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetUserWithPosts يقوم بإرجاع مستخدم مع علاقته (Posts) باستخدام Preload
func GetUserWithPosts(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := goeloquent.DB.Preload("Posts").First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "المستخدم غير موجود"})
		return
	}
	c.JSON(http.StatusOK, user)
}
