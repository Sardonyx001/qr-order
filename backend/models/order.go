package models

type Order struct {
	UUIDBaseModel
	Quantity     uint   `json:"quantity" gorm:"default:1;"`
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	TableID      string `json:"table_id"`
	Restaurant   Restaurant
	Customer     Customer
	Table        Table
}
