package handlers

import (
	"github.com/gin-gonic/gin"
	"it-courses/database"
	"it-courses/models"
	"net/http"
	"strconv"
)

func GetInstructors(c *gin.Context) {
	var instructors []models.Instructor
	if err := database.DB.Find(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении преподавателей"})
		return
	}
	c.JSON(http.StatusOK, instructors)
}

func AddInstructor(c *gin.Context) {
	var newInstructor models.Instructor
	if err := c.ShouldBindJSON(&newInstructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newInstructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении преподавателя"})
		return
	}

	c.JSON(http.StatusCreated, newInstructor)
}

func UpdateInstructor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var instructor models.Instructor

	if err := database.DB.First(&instructor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Преподаватель не найден"})
		return
	}

	if err := c.ShouldBindJSON(&instructor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении преподавателя"})
		return
	}

	c.JSON(http.StatusOK, instructor)
}

func DeleteInstructor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Delete(&models.Instructor{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении преподавателя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Преподаватель удалён"})
}
