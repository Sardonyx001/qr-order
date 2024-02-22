package models

type Order struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	Restaurant   Restaurant
	TableID      string `json:"table_id"`
	Table        Table
	ItemID       string `json:"item_id"`
	Item         Item
	Quantity     int    `json:"quantity"`
	Status       bool   `json:"status"`
	Options      string `json:"options"`
}
