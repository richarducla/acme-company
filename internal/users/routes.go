package users

import "github.com/labstack/echo/v4"

const (
	registerPath = "/auth/register"
	signInPath   = "/auth/login"
	signOutPath  = "/auth/logout"
)

func SetupAuthRoutes(h Handler, g *echo.Group) {
	g.POST(registerPath, h.SignUp)
	g.POST(signInPath, h.SingIn)
	g.POST(signOutPath, h.SingOut)
}
