package handlers

import "backend/services"

type Handlers struct {
	UserHandler
	AdminHandler
	AuthHandler
	RestaurantHandler
	CategoryHandler
	ItemHandler
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
			r: s.Restaurant,
			u: s.User,
			c: s.Category},
		CategoryHandler: &categoryHandler{s.Category},
		ItemHandler:     &itemHandler{s.Item},
	}
}
