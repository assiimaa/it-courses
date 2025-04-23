package handlers

import (
	"fmt"
	"it-courses/database"
	"it-courses/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.PaidAt = time.Now()

	if err := database.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment recorded", "payment": payment})
}

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	db := database.DB

	// Фильтрация
	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if courseID := c.Query("course_id"); courseID != "" {
		db = db.Where("course_id = ?", courseID)
	}

	// Пагинация
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	var pageInt, limitInt int
	fmt.Sscanf(page, "%d", &pageInt)
	fmt.Sscanf(limit, "%d", &limitInt)

	if pageInt <= 0 {
		pageInt = 1
	}
	offset := (pageInt - 1) * limitInt

	db = db.Limit(limitInt).Offset(offset)

	// Получение данных
	if err := db.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payments": payments})
}
