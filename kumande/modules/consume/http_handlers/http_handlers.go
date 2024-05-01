package httphandlers

import (
	"kumande/modules/consume/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetConsumeBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetConsumeBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
