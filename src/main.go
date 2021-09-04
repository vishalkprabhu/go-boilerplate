package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vishalkprabhu/go-boilerplate/controller"
	"github.com/vishalkprabhu/go-boilerplate/storage"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
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
