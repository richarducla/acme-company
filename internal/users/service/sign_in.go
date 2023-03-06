package service

import (
	"acme/pkg/errors"
	"acme/utils"

	"github.com/labstack/gommon/log"
)

type SignInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s service) SignIn(input SignInInput) (*string, *errors.HandlerError) {
	user, err := s.userRepository.Get(input.Username)
	if err != nil {
		log.Error("error getting user %w", err)
		return nil, errors.BadRequestErr
	}

	if err := utils.VerifyPassword(user.Password, input.Password); err != nil {
		log.Error("wrong username or password :%w ", err)
		return nil, errors.UnauthorizedErr
	}

	return &user.Username, nil
}
