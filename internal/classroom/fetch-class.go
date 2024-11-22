package classroom

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"

	"github.com/gin-gonic/gin"
)

func FetchClassForUser(c *gin.Context) {
	// Get the user ID from the cookie (or however it is stored in your application)
	userID, err := c.Cookie("user_id")

	if err != nil {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	db := config.GetDB()

	// Create a slice to store the classrooms
	var classrooms []models.Classroom
	result := db.Where("user_id = ?", userID).Find(&classrooms)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch classrooms"})
		return
	}

	c.JSON(200, gin.H{"classrooms": classrooms})
}

func FetchClassBySubject(c *gin.Context){
	// Retrieve from body in JSON
	var requestBody struct {
		Subject string `json:"subject"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	db := config.GetDB()

	// Create a slice to store the classrooms
	var classrooms []models.Classroom
	result := db.Where("subject = ?", requestBody.Subject).Find(&classrooms)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch classrooms"})
		return
	}

	c.JSON(200, gin.H{"classrooms": classrooms})
}