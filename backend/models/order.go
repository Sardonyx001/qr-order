package models

type Order struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
	TableID      string `json:"table_id"`
	Restaurant   Restaurant
	Customer     Customer
	Table        Table
}
