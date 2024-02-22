package models

type Customer struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	Restaurant   Restaurant
	TableID      string `json:"table_id"`
	Table        Table
	SessionID    string `json:"session_id"`
}
