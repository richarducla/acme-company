package users

import (
	usersService "acme/internal/users/service"
	"acme/pkg/errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler interface {
	SignUp(ec echo.Context) error
	SingIn(ec echo.Context) error
	SingOut(ec echo.Context) error
}

type handler struct {
	usersService usersService.Service
}

var _ Handler = &handler{}

func NewHandler(usersService usersService.Service) Handler {
	return &handler{
		usersService: usersService,
	}
}

// SignUp godoc
//
//	@Summary		Register an user
//	@Description	This endpoint save an user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user body		service.UserInput	true	"Add user"
//	@Success		201	{object}	RegisterResponse
//	@Failure		400	{object}	errors.HandlerError
//	@Failure		500	{object}	errors.HandlerError
//	@Router			/auth/register [post]
func (h handler) SignUp(ec echo.Context) error {
	var body usersService.UserInput

	if err := ec.Bind(&body); err != nil {
		log.Error("error in the bind %w", err)
		return ec.JSON(http.StatusBadRequest, errors.BadRequestErr)
	}

	user, err := h.usersService.Create(body)
	if err != nil {
		return ec.JSON(err.StatusCode, err)
	}

	return ec.JSON(http.StatusCreated, toRegisterResponse(user))
}

// SignIn godoc
//
//	@Summary		Login User
//	@Description	This endpoint singin an user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user body		service.SignInInput		true	"Login user"
//	@Success		200	{object}	Message
//	@Failure		400	{object}	errors.HandlerError
//	@Failure		401	{object}	errors.HandlerError
//	@Failure		500	{object}	errors.HandlerError
//	@Router			/auth/login [post]
func (h handler) SingIn(ec echo.Context) error {
	var credentials usersService.SignInInput

	if err := ec.Bind(&credentials); err != nil {
		log.Error("error in the bind %w", err)
		return ec.JSON(http.StatusBadRequest, errors.BadRequestErr)
	}

	username, err := h.usersService.SignIn(credentials)
	if err != nil {
		return ec.JSON(err.StatusCode, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = fmt.Sprintf("logged_in_%s", *username)
	cookie.Value = *username
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ec.SetCookie(cookie)

	return ec.JSON(http.StatusOK, Message{
		Message: statusSuccess,
	})
}

// SignOut godoc
//
//	@Summary		Logout user
//	@Description	This endpoint logout a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user body		service.SignInInput		true	"Logout user"
//	@Success		200	{object}	Message
//	@Failure		422	{object}	Message
//	@Router			/auth/logout [post]
func (h handler) SingOut(ec echo.Context) error {
	var credentials usersService.SignInInput

	if err := ec.Bind(&credentials); err != nil {
		log.Error("error in the bind %w", err)
		return ec.JSON(http.StatusBadRequest, err)
	}

	cookie, err := ec.Cookie(fmt.Sprintf("logged_in_%s", credentials.Username))
	if err != nil {
		return ec.JSON(http.StatusUnprocessableEntity, Message{
			Message: "user no logged",
		})
	}

	cookie.MaxAge = -1
	ec.SetCookie(cookie)
	return ec.JSON(http.StatusOK, Message{
		Message: statusLogout,
	})
}
