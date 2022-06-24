package books

import (
	"api/basic-rest-api/pkg/common/models"
	"api/basic-rest-api/pkg/common/requestbody"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddBook(c *gin.Context) {
	body := requestbody.BookRequestBody{}
	
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	result := h.DB.Create(&book)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}