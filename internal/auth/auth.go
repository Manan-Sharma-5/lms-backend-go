package auth

import (
	config "backend/internal/db-config"
	models "backend/pkg/model"
	"net/http"
	"time"

	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func validatePassword(password string) error {
	if len(password) < 8 || len(password) > 15 {
		return errors.New("password must be between 8 and 15 characters long")
	}

	// Must contain at least one digit
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return errors.New("password must contain at least one digit")
	}

	// Must contain at least one uppercase letter
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Must contain at least one lowercase letter
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Must contain at least one special character
	if match, _ := regexp.MatchString(`[!@#\$%\^&\*\(\)\-\+=]`, password); !match {
		return errors.New("password must contain at least one special character (!@#$%^&*()-+=)")
	}

	// Must not contain whitespace
	if match, _ := regexp.MatchString(`\s`, password); match {
		return errors.New("password must not contain any white spaces")
	}

	return nil
}


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

    if err := validatePassword(req.Password); err != nil {
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

	c.SetCookie("user_id", user.ID, 3600, "/", "localhost", false, true) // Adjust cookie parameters as needed
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

	c.SetCookie("user_id", user.ID, 3600, "/", "localhost", false, true) // Adjust cookie parameters as needed
    c.JSON(http.StatusOK, user)
}
