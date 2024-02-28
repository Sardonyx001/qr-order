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
}

func New(stores *stores.Stores, config *config.Config) *Services {
	return &Services{
		Auth:       &authService{Stores: stores, Config: config},
		User:       &userService{stores: stores},
		Admin:      &adminService{stores: stores},
		Restaurant: &restaurantService{stores: stores},
	}
}
