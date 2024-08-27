package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all openings",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST Opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT Opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}
