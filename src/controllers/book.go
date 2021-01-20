package controllers

import (
	r "goapi/ApiHelpers"
	"goapi/models"

	"github.com/gin-gonic/gin"
)

// ListBook returns BookModel JSON
func ListBook(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBook(&book)
	if err != nil {
		r.JSON(c, 404, book)
	} else {
		r.JSON(c, 200, book)
	}
}