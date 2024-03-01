package models

type Customer struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	TableID      string `json:"table_id"`
	SessionID    string `json:"session_id"`
}
