package reminder

import (
	"errors"
	"kumande/configs"
	"kumande/modules/stats"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReminderController struct {
	ReminderService ReminderService
	StatsService    stats.StatsService
}

func NewReminderController(weatherService ReminderService, statsService stats.StatsService) *ReminderController {
	return &ReminderController{
		ReminderService: weatherService,
		StatsService:    statsService,
	}
}

// Query
func (c *ReminderController) GetMostContextReminder(ctx *gin.Context) {
	// Param
	targetCol := ctx.Param("target_col")

	// Validator : Target Column Validator
	if !utils.Contains(configs.StatsReminderField, targetCol) {
		utils.BuildResponseMessage(ctx, "failed", "reminder", "target_col is not valid", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "reminder", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	reminder, err := c.StatsService.GetMostUsedContext("reminders", targetCol, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "reminder", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "reminder", "get", http.StatusOK, reminder, nil)
}
