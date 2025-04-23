package models

import "time"

type Schedule struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	CourseID     uint       `json:"course_id"`
	Course       Course     `json:"course"`
	InstructorID uint       `json:"instructor_id"`
	Instructor   Instructor `json:"instructor"`
	DayOfWeek    string     `json:"day_of_week"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      time.Time  `json:"end_time"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
