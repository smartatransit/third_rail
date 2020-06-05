package models

import "github.com/jinzhu/gorm"

type FeedbackType struct {
	gorm.Model
	Description string `gorm:"not null"`
}
