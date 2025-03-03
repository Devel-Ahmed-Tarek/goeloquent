package controllers

import (
	"net/http"

	"github.com/Devel-Ahmed-Tarek/goeloquent/app/models"
	"github.com/Devel-Ahmed-Tarek/goeloquent/goeloquent"
	"github.com/gin-gonic/gin"
)

// GetProducts يقوم بإرجاع جميع المنتجات مع دعم Pagination
func GetProducts(c *gin.Context) {
	var products []models.Product
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	result, err := goeloquent.Paginate(goeloquent.DB, &models.Product{}, &products, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
