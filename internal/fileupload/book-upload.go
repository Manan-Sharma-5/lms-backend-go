package file_upload

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"log"

	"github.com/gin-gonic/gin"
)

func BookRequest(c *gin.Context) {
    // Retrieve query parameters for file path and name
    var RequestBody struct {
        Title string `json:"title"`
        Author string `json:"author"`
        Description string `json:"description"`
        Price float64 `json:"price"`
    }

    // Get the user ID from the cookie (or however it is stored in your application)
    userID, err := c.Cookie("user_id")
    if err != nil {
        c.JSON(401, gin.H{"error": "User not authenticated"})
        return
    }

    // Parse the request body

    if err := c.BindJSON(&RequestBody); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request body"})
        return
    }
    // Create an AWS session
   
    db := config.GetDB()
    record := models.Book{
        Title: RequestBody.Title,
        Author: RequestBody.Author,
        Description: RequestBody.Description,
        Price: RequestBody.Price,
        Status: "pending",
        UserID: userID,
    }
    if err := db.Create(&record).Error; err != nil {
        log.Println("Failed to store file record:", err)
        c.JSON(500, gin.H{"error": "Failed to store file record in database"})
        return
    }

    // Return the PUT URL for uploading to the client
    c.JSON(200, gin.H{
        "message": "Successfully created book request",
        "body": RequestBody,
    })
}
