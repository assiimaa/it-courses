package handlers

import (
	"it-courses/database"
	"it-courses/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSchedule создаёт новое расписание
func CreateSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание записи
	schedule.CreatedAt = time.Now()
	if err := database.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule created", "schedule": schedule})
}

// GetSchedule получает все расписания с фильтрами и пагинацией
func GetSchedule(c *gin.Context) {
	var schedules []models.Schedule
	db := database.DB

	// Фильтрация по course_id
	if courseID := c.Query("course_id"); courseID != "" {
		db = db.Where("course_id = ?", courseID)
	}

	// Пагинация
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	offset, limit := getPagination(page, pageSize)

	db = db.Offset(offset).Limit(limit)

	if err := db.Preload("Course").Preload("Instructor").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schedule": schedules})
}

// UpdateSchedule обновляет расписание
func UpdateSchedule(c *gin.Context) {
	var schedule models.Schedule
	scheduleID := c.Param("id")

	if err := database.DB.First(&schedule, scheduleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	// Обновление записи
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated", "schedule": schedule})
}

// DeleteSchedule удаляет расписание
func DeleteSchedule(c *gin.Context) {
	var schedule models.Schedule
	scheduleID := c.Param("id")

	if err := database.DB.First(&schedule, scheduleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	if err := database.DB.Delete(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted"})
}

// Функция для пагинации
func getPagination(page string, pageSize string) (int, int) {
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	offset := (pageInt - 1) * pageSizeInt
	return offset, pageSizeInt
}
