package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Price        float64    `json:"price"`
	CategoryID   uint       `json:"category_id"`
	Category     Category   `json:"category" gorm:"foreignKey:CategoryID"`
	InstructorID uint       `json:"instructor_id"`
	Instructor   Instructor `json:"instructor" gorm:"foreignKey:InstructorID"`
}
