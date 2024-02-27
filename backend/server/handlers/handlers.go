package handlers

import "backend/services"

type Handlers struct {
	UserHanlder
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHanlder: &userHandler{s.User},
	}
}
