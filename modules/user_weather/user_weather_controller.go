package userWeather

import (
	"errors"
	"kumande/configs"
	"kumande/modules/stats"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserWeatherController struct {
	UserWeatherService UserWeatherService
	StatsService       stats.StatsService
}

func NewUserWeatherController(weatherService UserWeatherService, statsService stats.StatsService) *UserWeatherController {
	return &UserWeatherController{
		UserWeatherService: weatherService,
		StatsService:       statsService,
	}
}

// Query
func (c *UserWeatherController) GetMostContextUserWeather(ctx *gin.Context) {
	// Param
	targetCol := ctx.Param("target_col")

	// Validator : Target Column Validator
	if !utils.Contains(configs.StatsWeatherField, targetCol) {
		utils.BuildResponseMessage(ctx, "failed", "user weather", "target_col is not valid", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "user weather", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	weather, err := c.StatsService.GetMostUsedContext("user_weathers", targetCol, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "user weather", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "user weather", "get", http.StatusOK, weather, nil)
}
