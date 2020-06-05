package models

import "github.com/jinzhu/gorm"

type Direction struct {
	gorm.Model
	Feedback []Feedback
	Lines    []Line `gorm:"many2many:line_directions"`
	Name     string `gorm:"not null"`
}
