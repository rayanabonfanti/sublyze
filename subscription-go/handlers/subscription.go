package handlers

import (
	"subscription/database"
	"subscription/models"

	"github.com/gin-gonic/gin"
)

func CreateSubscription(c *gin.Context) {
	var sub models.Subscription

	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	db := database.GetDB()
	if err := db.Create(&sub).Error; err != nil {
		c.JSON(500, gin.H{"error": "Falha ao criar assinatura"})
		return
	}

	c.JSON(201, sub)
}

func GetAllSubscriptions(c *gin.Context) {
	var subs []models.Subscription
	db := database.GetDB()

	if err := db.Find(&subs).Error; err != nil {
		c.JSON(500, gin.H{"error": "Falha ao buscar assinaturas"})
		return
	}

	c.JSON(200, subs)
}
