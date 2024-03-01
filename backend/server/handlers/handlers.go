package handlers

import "backend/services"

type Handlers struct {
	UserHandler
	AdminHandler
	AuthHandler
	RestaurantHandler
	CategoryHandler
	ItemHandler
	TableHandler
	CustomerHandler
	OrderHandler
	OrderItemHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:  &userHandler{s.User},
		AdminHandler: &adminHandler{s.Admin},
		AuthHandler: &authHandler{
			s.Auth,
			s.User,
			s.Admin},
		RestaurantHandler: &restaurantHandler{
			s.Restaurant,
			s.User,
			s.Category},
		CategoryHandler: &categoryHandler{s.Category, s.Restaurant},
		TableHandler:    &tableHandler{s.Table},
		ItemHandler: &itemHandler{
			s.Item,
			s.Category,
			s.Restaurant},
		CustomerHandler:  &customerHandler{s.Customer},
		OrderHandler:     &orderHandler{s.Order, s.OrderItem},
		OrderItemHandler: &orderItemHandler{s.OrderItem},
	}
}
