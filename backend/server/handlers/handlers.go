package handlers

import "backend/services"

type Handlers struct {
	UserHandler
	AdminHandler
	AuthHandler
	RestaurantHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:  &userHandler{s.User},
		AdminHandler: &adminHandler{s.Admin},
		AuthHandler:  &authHandler{s.Auth},
		RestaurantHandler: &restaurantHandler{
			r: s.Restaurant,
			u: s.User},
	}
}
