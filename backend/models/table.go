package models

type Table struct {
	UUIDBaseModel
	RestaurantID string
	Restaurant   Restaurant
	Name         string `json:"name" gorm:"not null"`
	Empty        bool   `json:"empty" gorm:"not null"`
	Orders       []Order
}
