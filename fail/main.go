package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/api/v1/hello", func(c echo.Context) error {
		// if rand.Float64() > 0.75 {
		// 	return c.String(http.StatusOK, "Hello, Canary! - you are lucky!\n")
		// }
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error"))
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
