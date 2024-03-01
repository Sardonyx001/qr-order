package services

import (
	"backend/config"
	"backend/stores"
)

type Services struct {
	Auth       AuthService
	User       UserService
	Admin      AdminService
	Restaurant RestaurantService
	Item       ItemService
	Category   CategoryService
	Order      OrderService
	Table      TableService
	OrderItem  OrderItemService
	Customer   CustomerService
}

func New(stores *stores.Stores, config *config.Config) *Services {
	return &Services{
		Auth:       &authService{Stores: stores, Config: config},
		User:       &userService{stores: stores},
		Admin:      &adminService{stores: stores},
		Restaurant: &restaurantService{stores: stores},
		Item:       &itemService{stores: stores},
		Category:   &categoryService{stores: stores},
		Order:      &orderService{stores: stores},
		Table:      &tableService{stores: stores},
		OrderItem:  &orderItemService{stores: stores},
		Customer:   &customerService{stores: stores},
	}
}
