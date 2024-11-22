package fetchrequests

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"

	"github.com/gin-gonic/gin"
)

func FetchNotes(c *gin.Context) {
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
	var notes []models.Note
	result := db.Where("year = ? AND subject = ? AND status != ?", requestBody.Year, requestBody.SubjectCode, "pending").Find(&notes)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch notes"})
		return
	}

	// we need to return only the note whose status is not pending

	for i, note := range notes {
		if note.Status == "pending" {
			notes = append(notes[:i], notes[i+1:]...)
		}
	}

	// Return the notes

	c.JSON(200, gin.H{"notes": notes})
}
