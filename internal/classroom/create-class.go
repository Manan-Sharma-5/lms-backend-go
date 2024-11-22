package classroom

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	// Get the user ID from the cookie (or however it is stored in your application)
	userID, err := c.Cookie("user_id")

	if err != nil {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	// Fetch name, stream, subject, year, and createdDate from JSON request body
	var requestBody struct {
		Name        string `json:"name"`
		Stream      string `json:"stream"`
		Subject     string `json:"subject"`
		Year        int    `json:"year"`
		CreatedDate int64  `json:"createdDate"` // Expecting a Unix timestamp in milliseconds
		URL  	 string `json:"url"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Convert the Unix timestamp (in milliseconds) to a time.Time object
	createdDate := time.UnixMilli(requestBody.CreatedDate)

	// Create a new classroom
	classroom := models.Classroom{
		Name:        requestBody.Name,
		Stream:      requestBody.Stream,
		Subject:     requestBody.Subject,
		URL: 	   requestBody.URL,
		Year:        requestBody.Year,
		CreatedDate: createdDate,
		UserID:      userID,

	}

	// Save the classroom to the database
	db := config.GetDB()
	result := db.Create(&classroom)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create classroom"})
		return
	}

	// Return the classroom ID
	c.JSON(200, gin.H{"classroomID": classroom.ID})
}
