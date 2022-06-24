package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

// RegisterRoutes registers all the routes
func RegisterRoutes( r *gin.Engine, db *gorm.DB)  {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/v1/books")
	routes.GET("/:id", h.GetBook)
	routes.GET("/", h.GetBooks)
	routes.POST("/", h.AddBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)


}