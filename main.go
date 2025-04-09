package main

import (
	config "backend/internal/db-config"
	routes "backend/pkg"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.Use(func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        if origin == "http://localhost:3000" { // Update with allowed origins
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        }
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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