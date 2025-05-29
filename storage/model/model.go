package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	// Task structure
	Task struct {
		gorm.Model

		// Base information
		Name        string `gorm:"not null;size:100" validate:"required,min=3,max=100" json:"name"`
		Description string `gorm:"size:255" json:"description"`

		// Time information
		StartTime time.Time `gorm:"not null" validate:"required" json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}
)
