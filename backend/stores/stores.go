package stores

import "gorm.io/gorm"

type Stores struct {
	DB         *gorm.DB
	User       UserStore
	Admin      AdminStore
	Restaurant RestaurantStore
	Item       ItemStore
	Category   CategoryStore
	Table      TableStore
	Order      OrderStore
	OrderItem  OrderItemStore
	Customer   CustomerStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:         db,
		User:       &userStore{db},
		Admin:      &adminStore{db},
		Restaurant: &restaurantStore{db},
		Item:       &itemStore{db},
		Category:   &categoryStore{db},
		Table:      &tableStore{db},
		Order:      &orderStore{db},
		OrderItem:  &orderItemStore{db},
		Customer:   &customerStore{db},
	}
}
