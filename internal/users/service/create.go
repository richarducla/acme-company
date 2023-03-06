package service

import (
	"acme/internal/models"
	"acme/pkg/errors"
	"acme/utils"
	"strings"

	"github.com/labstack/gommon/log"
)

type UserInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
}

func (s service) Create(input UserInput) (models.User, *errors.HandlerError) {

	password, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Error("error encrypted password %w ", err)
		return models.User{}, errors.UnprocessableEntityErr.WithDescription(err.Error())
	}

	user := models.User{
		Username: input.Username,
		Password: password,
		Email:    strings.ToLower(input.Email),
	}

	user, err = s.userRepository.Create(user)
	if err != nil {
		return models.User{}, errors.UnprocessableEntityErr.WithDescription(err.Error())
	}

	return user, nil
}
