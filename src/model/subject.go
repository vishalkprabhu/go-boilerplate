package model

import (
	"time"
)

type Subject struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
