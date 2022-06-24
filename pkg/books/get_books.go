package books

import (
	"api/basic-rest-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	result := h.DB.Find(&books)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}