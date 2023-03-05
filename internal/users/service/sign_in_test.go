package service

import (
	"acme/internal/models"
	"acme/pkg/errors"
	"acme/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_SingIn(t *testing.T) {
	input := SignInInput{
		Username: "richard",
		Password: "1234",
	}
	t.Run("Should return error if get in repository fail", func(t *testing.T) {
		userService, userRepository := givenUserService()

		userRepository.On("Get", mock.Anything).Return(models.User{}, errors.InternalServerErrorErr)

		_, err := userService.SignIn(input)
		assert.Error(t, err)
		assert.Equal(t, err.Message, BadRequestErrorMsg)
		assert.Equal(t, err.StatusCode, http.StatusBadRequest)
	})

	t.Run("Should return error for the password is wrong", func(t *testing.T) {
		password, _ := utils.HashPassword("12345")
		userService, userRepository := givenUserService()

		userRepository.On("Get", mock.Anything).Return(models.User{
			ID:       1,
			Username: input.Username,
			Email:    "rj@test.com",
			Password: password,
		}, nil)

		_, err := userService.SignIn(input)
		assert.Error(t, err)
		assert.Equal(t, err.Message, UnauthorizedErrMsg)
		assert.Equal(t, err.StatusCode, http.StatusUnauthorized)
	})

	t.Run("Should return username if password and user is correct", func(t *testing.T) {
		password, _ := utils.HashPassword(input.Password)
		userService, userRepository := givenUserService()

		userRepository.On("Get", mock.Anything).Return(models.User{
			ID:       1,
			Username: input.Username,
			Email:    "rj@test.com",
			Password: password,
		}, nil)

		username, err := userService.SignIn(input)
		assert.Nil(t, err)
		assert.Equal(t, "richard", *username)
	})
}
