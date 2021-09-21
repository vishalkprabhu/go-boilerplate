package model

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Class string `json:"class"`
	Order uint   `json:"order"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
