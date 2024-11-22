package fetchrequests

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"

	"github.com/gin-gonic/gin"
)

func ViewBooks(c *gin.Context){
	db := config.GetDB()

	var books []models.Book

	result := db.Find(&books)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(200, gin.H{"books": books})
}