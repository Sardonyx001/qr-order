package models

type Subscription struct {
	UUIDBaseModel
	Paid_year_month string `json:"paid_year_month"`
	RestaurantID    string `json:"restaurantId"`
	Restaurant      Restaurant
}