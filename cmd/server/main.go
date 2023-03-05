// @title API Acme
// @version 1.0.0
// @description this is the api for the auth users in ACME company

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/acme
package main

import (
	"acme/cmd/config"
	"acme/internal/users"
	umethods "acme/internal/users/repository/methods"
	uservice "acme/internal/users/service"

	"fmt"
	"net/http"

	_ "acme/docs"
	"acme/pkg/postgres"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

const (
	prefixPath = "/api/acme"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Errorf("can't read config: %w", err))
	}

	db := createDB(cfg)

	server := setupServer(cfg)

	router := server.Group(prefixPath)
	setupRoutes(router, db)

	server.Logger.Fatal(server.Start(cfg.SrvAddr()))
}

func setupServer(c config.Config) *echo.Echo {
	server := echo.New()

	server.GET(prefixPath+"/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	if c.Env == "dev" {
		serveDocs(server)
	}

	return server
}

func createDB(c config.Config) *gorm.DB {
	db, err := postgres.Init(c)
	if err != nil {
		panic(err)
	}
	return db
}

func setupRoutes(router *echo.Group, db *gorm.DB) {
	userRepository := umethods.NewRepository(db)
	userService := uservice.NewService(userRepository)
	userHandler := users.NewHandler(userService)

	users.SetupAuthRoutes(userHandler, router)
}

func serveDocs(server *echo.Echo) {
	server.GET(prefixPath+"/docs/*", echoSwagger.WrapHandler)
}
