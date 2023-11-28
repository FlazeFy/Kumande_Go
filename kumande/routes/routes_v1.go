package routes

import (
	"net/http"

	syshandlers "kumande/modules/stats/http_handlers"

	"github.com/labstack/echo"
)

// NON ORM
func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Kumande")
	})

	// =============== Public routes ===============
	// Dictionary
	e.GET("api/v1/stats/consume_from/:ord/:limit", syshandlers.GetTotalConsumeByFrom)
	e.GET("api/v1/stats/consume_type/:ord/:limit", syshandlers.GetTotalConsumeByType)
	e.GET("api/v1/stats/consume_ing/:ord/:limit", syshandlers.GetTotalConsumeByMainIng)
	e.GET("api/v1/stats/consume_prov/:ord/:limit", syshandlers.GetTotalConsumeByProvide)

	// =============== Private routes ===============

	return e
}
