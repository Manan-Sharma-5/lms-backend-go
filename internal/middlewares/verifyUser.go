package middleware

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Retrieve the cookie
        userID, err := c.Cookie("user_id")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no cookie found"})
            c.Abort()
            return
        }

        // Check if the user exists in the database
        var user models.User
        db := config.GetDB()
        if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, user not found"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            }
            c.Abort()
            return
        }

        // Pass the user data through context
        c.Set("user", user)

        // Continue with the next handler
        c.Next()
    }
}
