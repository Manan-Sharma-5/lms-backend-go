package fetchrequests

import (
	config "backend/internal/db-config"

	"github.com/gin-gonic/gin"
)

func FetchSubjectsForNotes(c *gin.Context) {
    year := c.Query("year")
    stream := c.Query("stream")
    if year == "" || stream == "" {
        c.JSON(400, gin.H{"error": "Year and Stream are required"})
        return
    }

    db := config.GetDB()

    var subjects []string

	err := db.Table("notes").Select("DISTINCT subject").Where("year = ? AND stream = ?", year, stream).Pluck("subject", &subjects).Error

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch subjects"})
        return
    }

    c.JSON(200, gin.H{"subjects": subjects})
}

func FetchSubjectsForPapers(c *gin.Context) {
    year := c.Query("year")
    stream := c.Query("stream")
    if year == "" || stream == "" {
        c.JSON(400, gin.H{"error": "Year and Stream are required"})
        return
    }

    db := config.GetDB()

    var subjects []string

    err := db.Table("previous_year_questions").Select("DISTINCT subject").Where("year = ? AND stream = ?", year, stream).Pluck("subject", &subjects).Error

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch subjects"})
        return
    }

    c.JSON(200, gin.H{"subjects": subjects})
}

func FetchSubjectsForClasses(c *gin.Context) {
   
    db := config.GetDB()

    var subjects []string

    err := db.Table("classrooms").Select("DISTINCT subject").Pluck("subject", &subjects).Error

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch subjects"})
        return
    }

    c.JSON(200, gin.H{"subjects": subjects})
}