package http_app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	server *echo.Echo
}

type ClassService interface {
	List() []string
	Create() bool
}

func Initialize(cs ClassService) *App {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ag := e.Group("/admin")

	// Admin Routes
	ag.GET("", hello)
	ag.GET("/classes/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, cs.List())
	})

	//default index
	e.GET("/", hello)
	// Start server

	return &App{
		server: e,
	}

}

func (app *App) Run() {
	app.server.Logger.Fatal(app.server.Start(":1200"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! this is test")
}
