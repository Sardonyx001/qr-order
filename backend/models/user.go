package models

// User Status field enum
type Status string

const (
	Active   Status = "ACTIVE"
	Inactive Status = "INACTIVE"
)

// User model
type User struct {
	UUIDBaseModel
	Username     string        `json:"username"    gorm:"unique; not null"`
	PasswordHash string        `json:"-"           gorm:"not null"`
	Status       Status        `json:"status"      gorm:"default:'ACTIVE';not null"`
	AdminID      string        `json:"-"`
	Admin        Admin         `json:"-"`
	Restaurants  []*Restaurant `json:"restaurants" gorm:"many2many:user_restaurants;"`
}
