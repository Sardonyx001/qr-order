package models

type OrderItem struct {
	UUIDBaseModel
	Quantity     uint   `json:"quantity" gorm:"default:1"`
	Status       bool   `json:"status" gorm:"default:false;"`
	Options      string `json:"options"`
	RestaurantID string `json:"restaurant_id"`
	OrderID      string `json:"order_id"`
	Restaurant   Restaurant
	Order        Order
}
