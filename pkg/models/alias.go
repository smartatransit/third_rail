package models

import (
	"time"
)

type Alias struct {
	ID               uint   `json:"-" gorm:"primary_key"`
	NamedElementType string `gorm:"not null"`
	NamedElementID   uint   `json:"-" gorm:"not null"`
	Alias            string
	Description      string
	CreatedAt        time.Time  `json:"-"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-" sql:"index"`
}

type Aliases []Alias

func (a Aliases) String(i int) string {
	return a[i].Alias
}

func (a Aliases) Len() int {
	return len(a)
}
