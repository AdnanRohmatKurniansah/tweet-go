package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/database"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Tweet-go API",
			"status": "It's works",
			"version": "1.0.0",
			"timestamp": time.Now(),
		})
	})

	routes.SetupRoutes(r, cfg, db)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.APP_PORT)
	r.Run(server)
}