package books

import (
	"api/basic-rest-api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.JSON(http.StatusOK, &book)
}