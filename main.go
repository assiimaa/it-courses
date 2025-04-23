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

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/courses", handlers.GetCourses)
		protected.POST("/courses", handlers.CreateCourse)
		protected.PUT("/courses/:id", handlers.UpdateCourse)
		protected.DELETE("/courses/:id", handlers.DeleteCourse)

		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)
		protected.PUT("/categories/:id", handlers.UpdateCategory)
		protected.DELETE("/categories/:id", handlers.DeleteCategory)

		protected.GET("/instructors", handlers.GetInstructors)
		protected.POST("/instructors", handlers.AddInstructor)
		protected.PUT("/instructors/:id", handlers.UpdateInstructor)
		protected.DELETE("/instructors/:id", handlers.DeleteInstructor)

		protected.GET("/payments", handlers.GetPayments)
		protected.POST("/payments", handlers.CreatePayment)

		protected.GET("/schedule", handlers.GetSchedule)
		protected.POST("/schedule", handlers.CreateSchedule)
		protected.PUT("/schedule/:id", handlers.UpdateSchedule)
		protected.DELETE("/schedule/:id", handlers.DeleteSchedule)

	}

	r.Run(":8080")
}
