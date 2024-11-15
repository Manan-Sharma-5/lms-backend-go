package auth

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SignUp registers a new user
func SignUp(c *gin.Context) {
    var req struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
        Role     string `json:"role" binding:"required"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password
    salt := 10
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), salt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }

    // Create user model
    user := models.User{
        Name:       req.Name,
        Email:      req.Email,
        Password:   string(hashedPassword),
        Role:       req.Role,
        DateJoined: time.Now(),
    }

    // Save user in the database
    if err := config.GetDB().Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

	c.SetCookie("user_id", user.ID, 3600, "/", "172.16.165.162", false, true) // Adjust cookie parameters as needed
    c.JSON(http.StatusCreated, user)
}

// SignIn logs in an existing user
func SignIn(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    db := config.GetDB()

    // Find user by email
    if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
        return
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
        return
    }

	c.SetCookie("user_id", user.ID, 3600, "/", "172.16.165.162", false, true) // Adjust cookie parameters as needed
    c.JSON(http.StatusOK, user)
}
