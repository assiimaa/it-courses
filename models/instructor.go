package models

import "gorm.io/gorm"

type Instructor struct {
	gorm.Model
	Name string `json:"name"`
}
