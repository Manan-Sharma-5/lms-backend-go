package file_upload

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// FileUpload generates a PUT URL for uploading and a GET URL for accessing the file.
func FileUpload(c *gin.Context) {
    region := "eu-north-1"
    bucket := "software-engineering-project-s3"

    // Retrieve query parameters for file path and name
    year := c.Query("year")
    subjectCode := c.Query("subjectCode")
    filename := c.Query("filename")
    if filename == "" {
        c.JSON(400, gin.H{"error": "Filename is required"})
        return
    }

    // Get the user ID from the cookie (or however it is stored in your application)
    userID, err := c.Cookie("user_id")
    if err != nil {
        c.JSON(401, gin.H{"error": "User not authenticated"})
        return
    }

    // Create an AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        log.Println("Failed to create AWS session:", err)
        c.JSON(500, gin.H{"error": "Failed to create AWS session"})
        return
    }

    // Create an S3 service client
    svc := s3.New(sess)

    // Generate a pre-signed PUT URL for the client to upload the file
    key := "notes/" + year + "/" + subjectCode + "/" + filename + ".pdf"
    putReq, _ := svc.PutObjectRequest(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
    putURL, err := putReq.Presign(15 * time.Minute)
    if err != nil {
        log.Println("Failed to sign PUT request:", err)
        c.JSON(500, gin.H{"error": "Failed to sign PUT request"})
        return
    }

    // Generate a pre-signed GET URL for accessing the uploaded file
    getReq, _ := svc.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
		getURL, err := getReq.Presign(7 * 24 * time.Hour) 
    if err != nil {
        log.Println("Failed to sign GET request:", err)
        c.JSON(500, gin.H{"error": "Failed to sign GET request"})
        return
    }

	yearInt, _ := strconv.Atoi(year)

    db := config.GetDB()
    record := models.Note{
		Title: 	 		filename,
		Year:			yearInt,
		Description: 	c.Query("description"),
		Subject: 		subjectCode,
		UserID: 		userID,
		UploadDate: 	time.Now(),
		Stream: 		c.Query("stream"),
		Status: 		"pending",
		Content: 		getURL,
    }
    if err := db.Create(&record).Error; err != nil {
        log.Println("Failed to store file record:", err)
        c.JSON(500, gin.H{"error": "Failed to store file record in database"})
        return
    }

    // Return the PUT URL for uploading to the client
    c.JSON(200, gin.H{
        "upload_url": putURL,
        "message": "File ready to be uploaded",
    })
}
