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
	result, err := repositories.GetTotalConsumeByFrom(page, 10, "api/v1/stats/consume_from/"+ord+"/"+limit, ord, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
