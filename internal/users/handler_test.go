package users

import (
	"acme/internal/models"
	"acme/internal/users/service"
	"acme/pkg/errors"
	"acme/test/mocks/services"
	"acme/test/setuptesting"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	endpoint := "/auth/register"
	userInput := service.UserInput{
		Username: "richard",
		Password: "1234",
		Email:    "rj@test.com",
	}
	t.Run("Should return bad request if body is invalid", func(t *testing.T) {
		badBody := `{asd}`

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, badBody, nil, nil, nil)

		h := &handler{}
		err := h.SignUp(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("Should return error for fail in service", func(t *testing.T) {

		body := setuptesting.BuildBody(userInput)

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, body, nil, nil, nil)

		userService := new(services.UserService)

		userService.On("Create", userInput).Return(models.User{}, errors.UnprocessableEntityErr.WithDescription("some error"))

		h := &handler{
			usersService: userService,
		}

		err := h.SignUp(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Result().StatusCode)
	})

	t.Run("Should return response successfully", func(t *testing.T) {

		body := setuptesting.BuildBody(userInput)

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, body, nil, nil, nil)

		userService := new(services.UserService)

		userService.On("Create", userInput).Return(models.User{
			Username: "richard",
			Email:    "rj@test.com",
		}, nil)

		h := &handler{
			usersService: userService,
		}

		err := h.SignUp(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)
	})
}

func TestHandler_SignIn(t *testing.T) {
	endpoint := "/auth/login"
	signInInput := service.SignInInput{
		Username: "richard",
		Password: "1234",
	}
	t.Run("Should return bad request if body is invalid", func(t *testing.T) {
		badBody := `{asd}`

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, badBody, nil, nil, nil)

		h := &handler{}
		err := h.SingIn(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("Should return error for fail in service", func(t *testing.T) {
		body := setuptesting.BuildBody(signInInput)

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, body, nil, nil, nil)

		userService := new(services.UserService)

		userService.On("SignIn", signInInput).Return(nil, errors.UnprocessableEntityErr.WithDescription("invalid credentials"))

		h := &handler{
			usersService: userService,
		}

		err := h.SingIn(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Result().StatusCode)
	})

	t.Run("Should return response successfully", func(t *testing.T) {
		response := ""
		body := setuptesting.BuildBody(signInInput)

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, body, nil, nil, nil)

		userService := new(services.UserService)

		userService.On("SignIn", signInInput).Return(&response, nil)

		h := &handler{
			usersService: userService,
		}

		err := h.SingIn(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})
}

func TestHandler_SignOut(t *testing.T) {
	endpoint := "/auth/logout"
	signInInput := service.SignInInput{
		Username: "richard",
		Password: "1234",
	}
	t.Run("Should return bad request if body is invalid", func(t *testing.T) {
		badBody := `{asd}`

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, badBody, nil, nil, nil)

		h := &handler{}
		err := h.SingOut(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("Should return error if user no logged", func(t *testing.T) {
		body := setuptesting.BuildBody(signInInput)

		_, _, rec, c := setuptesting.ServerTest(http.MethodPost, endpoint, body, nil, nil, nil)

		userService := new(services.UserService)

		h := &handler{
			usersService: userService,
		}

		err := h.SingOut(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Result().StatusCode)
	})

	t.Run("Should logout successfully", func(t *testing.T) {
		body := setuptesting.BuildBody(signInInput)

		cookieName := fmt.Sprintf("logged_in_%s", signInInput.Username)
		_, _, rec, c := setuptesting.ServerTest(
			http.MethodPost,
			endpoint,
			body,
			nil,
			nil,
			map[string]string{
				cookieName: signInInput.Username,
			},
		)

		userService := new(services.UserService)

		h := &handler{
			usersService: userService,
		}

		err := h.SingOut(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})
}
