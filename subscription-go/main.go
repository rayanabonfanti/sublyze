package main

import (
	"log"
	"subscription/database"
	"subscription/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.POST("/subscriptions", handlers.CreateSubscription)
	r.GET("/subscriptions", handlers.GetAllSubscriptions)

	log.Println("Server running on port 8081")
	r.Run(":8081")
}
