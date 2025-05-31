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
		Name        string `gorm:"not null;size:100"`
		Description string `gorm:"size:255"`

		// Time information
		StartTime time.Time `gorm:"not null"`
		EndTime   time.Time // UNUSED
	}
)
