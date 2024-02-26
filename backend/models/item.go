package models

type Item struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	Restaurant   Restaurant
	CategoryID   string `json:"category_id"`
	Category     Category
	Name         string `json:"name"`
	Text         string `json:"text"`
	Price        int    `json:"price"`
	Options      string `json:"options"`
	Img          string `json:"img"`
}
