package domain

import "gorm.io/gorm"

type Job struct {
	gorm.Model

	Title       string  `json:"title" gorm:"size:255;not null"`
	Description string  `json:"description" gorm:"size:255;not null"`
	Company     string  `json:"company" gorm:"size:255;not null"`
	Location    string  `json:"location" gorm:"size:255;not null"`
	Salary      float64 `json:"salary" gorm:"not null"`
}
