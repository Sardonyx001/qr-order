package models

// Admin model
type Admin struct {
	UUIDBaseModel
	Username     string `json:"name" gorm:"unique;not null"`
	PasswordHash string `json:"-"    gorm:"not null"`
	Users        []User `json:"users"`
}
