package main

import (
	config "backend/internal/db-config"
	routes "backend/pkg"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
    r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.Routes(r)

	dbConfig := config.Config{
        Host:     "localhost",
        Port:     5432,
        User:     "user",
        Password: "password",
        DBName:   "collab",
        SSLMode:  "disable", // or "enable" as needed
    }

    // Initialize the database connection
    config.InitializeDB(dbConfig)

	r.Run(":8080")
}