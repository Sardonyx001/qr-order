package models

import "gorm.io/gorm"

// User Status field enum
type Status string

const (
	Active   Status = "ACTIVE"
	Inactive Status = "INACTIVE"
)

// User model
type User struct {
	gorm.Model
	Username     string `json:"name"   gorm:"unique; not null"`
	PasswordHash string `json:"-"      gorm:"not null"`
	Status       Status `json:"status" gorm:"not null"`
	Restaurants  []Restaurant
}
