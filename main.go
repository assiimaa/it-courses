package main

import (
	"github.com/gin-gonic/gin"
	"it-courses/database"
	"it-courses/handlers"
	"it-courses/middleware"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	// Публичные маршруты
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	// Защищённые маршруты (JWT)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Courses
		protected.GET("/courses", handlers.GetCourses)
		protected.POST("/courses", handlers.CreateCourse)
		protected.PUT("/courses/:id", handlers.UpdateCourse)
		protected.DELETE("/courses/:id", handlers.DeleteCourse)

		// Categories
		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)
		protected.PUT("/categories/:id", handlers.UpdateCategory)
		protected.DELETE("/categories/:id", handlers.DeleteCategory)

		// Instructors
		protected.GET("/instructors", handlers.GetInstructors)
		protected.POST("/instructors", handlers.AddInstructor)
		protected.PUT("/instructors/:id", handlers.UpdateInstructor)
		protected.DELETE("/instructors/:id", handlers.DeleteInstructor)
	}

	r.Run(":8080")
}
