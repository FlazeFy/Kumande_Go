package httphandlers

import (
	"kumande/modules/stats/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetTotalConsumeByFrom(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "consume_from"

	result, err := repositories.GetTotalConsumeMulti(page, 10, "api/v1/stats/consume_from/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalConsumeByType(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	limit := c.Param("limit")
	view := "consume_type"

	result, err := repositories.GetTotalConsumeMulti(page, 10, "api/v1/stats/consume_type/"+ord+"/"+limit, ord, limit, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
