package httphandlers

import (
	"kumande/modules/stats/repositories"
	"net/http"

	"github.com/labstack/echo"
)

func GetTotalConsumeByFrom(c echo.Context) error {
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "consume_from"

	result, err := repositories.GetTotalConsumeMulti("api/v1/stats/consume_from/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalConsumeByType(c echo.Context) error {
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "consume_type"

	result, err := repositories.GetTotalConsumeMulti("api/v1/stats/consume_type/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalConsumeByMainIng(c echo.Context) error {
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "main_ing"

	result, err := repositories.GetTotalConsumeByConsumeDetail("api/v1/stats/consume_ing/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalConsumeByProvide(c echo.Context) error {
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "provide"

	result, err := repositories.GetTotalConsumeByConsumeDetail("api/v1/stats/consume_prov/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
