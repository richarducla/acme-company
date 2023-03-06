package service

import (
	"acme/internal/models"
	"acme/internal/users/domain"
	"acme/pkg/errors"
)

//go:generate mockery --name=Service --filename=userservice.go --structname=UserService --output=../../../test/mocks/services --outpkg=services

type Service interface {
	Create(input UserInput) (models.User, *errors.HandlerError)
	SignIn(input SignInInput) (*string, *errors.HandlerError)
}

type service struct {
	userRepository domain.Repository
}

func NewService(userRepository domain.Repository) Service {
	return service{userRepository: userRepository}
}
