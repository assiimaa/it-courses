package handlers

import (
	"it-courses/database"
	"it-courses/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	instructorIDStr := c.Query("instructor_id")

	pageStr := c.DefaultQuery("page", "1")    // по умолчанию страница 1
	limitStr := c.DefaultQuery("limit", "10") // по умолчанию 10 элементов на страницу

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var courses []models.Course
	query := database.DB.Preload("Category").Preload("Instructor")

	if categoryIDStr != "" {
		query = query.Where("category_id = ?", categoryIDStr)
	}
	if instructorIDStr != "" {
		query = query.Where("instructor_id = ?", instructorIDStr)
	}

	if err := query.Limit(limit).Offset(offset).Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении курсов"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCourseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID курса"})
		return
	}

	var course models.Course
	if err := database.DB.Preload("Category").Preload("Instructor").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Курс не найден"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var newCourse models.Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newCourse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании курса"})
		return
	}
	c.JSON(http.StatusCreated, newCourse)
}

func UpdateCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID курса"})
		return
	}

	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Курс не найден"})
		return
	}

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении курса"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID курса"})
		return
	}

	if err := database.DB.Delete(&models.Course{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении курса"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Курс удалён"})
}
