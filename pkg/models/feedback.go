package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Feedback struct {
	gorm.Model
	StationID   int
	LineID      int
	DirectionID int
	SourceID    int
	Source      FeedbackSource
	TypeID      int
	Type        FeedbackType
	Description string `gorm:"not null"`
	ThumbsUp    int
	ThumbsDown  int
	ExpiresAt   time.Time
}
