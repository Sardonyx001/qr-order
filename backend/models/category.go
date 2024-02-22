package models

type Category struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	Restaurant   Restaurant
	Name         string `json:"name" gorm:"unique; not null"`
	Description  string `json:"description"`
}
