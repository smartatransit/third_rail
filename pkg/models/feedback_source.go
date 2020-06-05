package models

import "github.com/jinzhu/gorm"

type FeedbackSource struct {
	gorm.Model
	SourceID   string
	SourceType string
}
