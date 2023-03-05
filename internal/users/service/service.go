package service

import (
	"acme/internal/users/domain"
)

type Service struct {
	userRepository domain.Repository
}

func NewService(userRepository domain.Repository) Service {
	return Service{userRepository: userRepository}
}
