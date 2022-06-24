package books

import (
	"api/basic-rest-api/pkg/common/models"
	"api/basic-rest-api/pkg/common/requestbody"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handler) UpdateBook(c *gin.Context) {
	body := requestbody.BookRequestBody{}
	id := c.Param("id")

	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	body.Title = book.Title
	body.Author = book.Author
	body.Description = book.Description

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}