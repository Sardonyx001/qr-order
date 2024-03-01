package models

type Item struct {
	UUIDBaseModel
	RestaurantID string `json:"restaurant_id"`
	CategoryID   string `json:"category_id"`
	Name         string `json:"name"`
	Text         string `json:"text"`
	Price        int    `json:"price"`
	Options      string `json:"options"`
	Img          string `json:"img"`
}
