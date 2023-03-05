package service

import (
	"acme/test/mocks/repositories"
)

const (
	InternalServerErrorMsg = "internal-server-error"
	BadRequestErrorMsg     = "bad-request-error"
	UnprocessableEntityMsg = "unprocessable-entity-error"
	UnauthorizedErrMsg     = "unauthorized-error"
)

func givenUserService() (Service, *repositories.UserRepository) {
	userRepository := new(repositories.UserRepository)

	userService := NewService(userRepository)

	return userService, userRepository
}
