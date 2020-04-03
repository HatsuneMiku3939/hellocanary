package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/api/v1/hello", func(c echo.Context) error {
		time.Sleep(4 * time.Second)
		return c.String(http.StatusOK, "Hello, Canary! - slow\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
