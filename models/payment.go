package models

import "time"

type Payment struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	UserID   uint      `json:"user_id"`
	CourseID uint      `json:"course_id"`
	Amount   float64   `json:"amount"`
	PaidAt   time.Time `json:"paid_at"`
}
