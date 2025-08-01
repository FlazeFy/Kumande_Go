package routes

import (
	"net/http"

	authhandlers "kumande/modules/auth/http_handlers"
	consumehandlers "kumande/modules/consume/http_handlers"
	syshandlers "kumande/modules/stats/http_handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NON ORM
func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Kumande")
	})

	// =============== Public routes ===============
	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/logout", authhandlers.SignOut)

	// Stats
	e.GET("api/v1/stats/consume_from/:ord/:limit", syshandlers.GetTotalConsumeByFrom)
	e.GET("api/v1/stats/consume_type/:ord/:limit", syshandlers.GetTotalConsumeByType)
	e.GET("api/v1/stats/consume_ing/:ord/:limit", syshandlers.GetTotalConsumeByMainIng)
	e.GET("api/v1/stats/consume_prov/:ord/:limit", syshandlers.GetTotalConsumeByProvide)
	e.GET("api/v1/stats/budget/:year", syshandlers.GetBudgetSpendingYear)
	e.GET("api/v1/stats/count/payment", syshandlers.GetSpendingInfo)
	e.GET("api/v1/stats/count/body", syshandlers.GetBodyInfo)

	// Consume
	e.GET("api/v1/consume/search/:slug", consumehandlers.GetConsumeBySlug)

	// =============== Private routes ===============

	return e
}
