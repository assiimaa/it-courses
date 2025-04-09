package main

import (
	"github.com/gin-gonic/gin"
	"it-courses/database"
	"it-courses/handlers"
)

func main() {

	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/courses", handlers.GetCourses)
	r.POST("/courses", handlers.CreateCourse)
	r.PUT("/courses/:id", handlers.UpdateCourse)
	r.DELETE("/courses/:id", handlers.DeleteCourse)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)
	r.PUT("/categories/:id", handlers.UpdateCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)

	r.GET("/instructors", handlers.GetInstructors)
	r.POST("/instructors", handlers.AddInstructor)
	r.PUT("/instructors/:id", handlers.UpdateInstructor)
	r.DELETE("/instructors/:id", handlers.DeleteInstructor)
	
	r.Run(":8080")
}
