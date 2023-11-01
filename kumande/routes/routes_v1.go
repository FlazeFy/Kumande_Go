package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// NON ORM
func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Kumande")
	})

	// =============== Public routes ===============

	// =============== Private routes ===============

	return e
}
