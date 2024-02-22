package models

// User model
type Admin struct {
	UUIDBaseModel
	Username     string `json:"name" gorm:"unique;not null"`
	PasswordHash string `json:"-"    gorm:"not null"`
}
