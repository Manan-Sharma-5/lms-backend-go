package fetchrequests

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"

	"github.com/gin-gonic/gin"
)

func FetchPYQS(c *gin.Context) {
	// Retrieve from body in JSON
	var requestBody struct {
		Year        int    `json:"year"`
		SubjectCode string `json:"subjectCode"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	db := config.GetDB()

	// Create a slice to store the notes
	var pyqs []models.PreviousYearQuestion
	result := db.Where("year = ? AND subject = ?", requestBody.Year, requestBody.SubjectCode).Find(&pyqs)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(200, gin.H{"pyqs": pyqs})
}
