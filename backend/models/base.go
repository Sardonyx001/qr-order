package models

import (
	"time"

	"gorm.io/gorm"
)

type UUIDBaseModel struct {
	ID        string `gorm:"primaryKey;size:255;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
