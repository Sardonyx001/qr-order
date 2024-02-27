package services

import "backend/stores"

type Services struct {
	User UserService
}

func New(stores *stores.Stores) *Services {
	return &Services{
		User: &userService{stores: stores},
	}
}
