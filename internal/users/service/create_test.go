package service

import (
	"acme/internal/models"
	"acme/pkg/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_Create(t *testing.T) {
	input := UserInput{
		Username: "richard",
		Password: "1234",
		Email:    "rj@test.com",
	}
	t.Run("Should return error if create in repository fail", func(t *testing.T) {
		userService, userRepository := givenUserService()

		userRepository.On("Create", mock.Anything).Return(models.User{}, errors.UnprocessableEntityErr)

		_, err := userService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, err.Message, UnprocessableEntityMsg)
		assert.Equal(t, err.StatusCode, http.StatusUnprocessableEntity)
	})

	t.Run("Should return user", func(t *testing.T) {
		userService, userRepository := givenUserService()

		userRepository.On("Create", mock.Anything).Return(models.User{
			ID:       1,
			Username: input.Username,
			Email:    input.Email,
		}, nil)

		user, err := userService.Create(input)
		assert.Nil(t, err)
		assert.IsType(t, models.User{}, user)
		assert.Equal(t, input.Email, user.Email)
	})
}
