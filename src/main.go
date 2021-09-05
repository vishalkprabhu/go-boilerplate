package main

import (
	"fmt"
	"net/http"
	"os"

	"go-boilerplate/src/controller"
	"go-boilerplate/src/settings"
	"go-boilerplate/src/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	env := os.Getenv("ENV_NAME")
	yamlPath := fmt.Sprintf("config/%s.yaml", env)

	config, errCfg := settings.GetSettings(yamlPath)

	if errCfg != nil {
		e.Logger.Fatal(errCfg)
	}

	var dbCfg storage.DbConfig
	dbCfg.User = config.Database.DbUser
	dbCfg.Password = config.Database.DbPass
	dbCfg.Host = config.Database.DbHost
	dbCfg.Password = config.Database.DbPass
	dbCfg.Port = config.Database.DbPort

	storage.NewDB(dbCfg)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/students", controller.GetStudents)

	// Start server
	e.Logger.Fatal(e.Start(":1200"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! this is test")
}
