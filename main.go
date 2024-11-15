package main

import (
	config "backend/internal/db-config"
	routes "backend/pkg"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routes.Routes(r)

	dbConfig := config.Config{
        Host:     "localhost",
        Port:     5432,
        User:     "user",
        Password: "user",
        DBName:   "collab",
        SSLMode:  "disable", // or "enable" as needed
    }

    // Initialize the database connection
    config.InitializeDB(dbConfig)

	r.Run(":8080")
}