package reminder

import (
	"errors"
	"kumande/configs"
	"kumande/modules/stats"
	"kumande/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// Command
func (c *ReminderController) HardDeleteReminderByID(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	reminderID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "reminder", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "reminder", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Reminder By ID
	err = c.ReminderService.HardDeleteReminderByID(reminderID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "reminder", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "reminder", "hard delete", http.StatusOK, nil, nil)
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

func (c *ReminderController) GetAllReminder(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "reminder", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Reminder
	res, total, err := c.ReminderService.GetAllReminder(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "reminder", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	metadata := gin.H{
		"total":       total,
		"page":        pagination.Page,
		"limit":       pagination.Limit,
		"total_pages": totalPages,
	}

	res = utils.StripFields(res, "created_by")
	utils.BuildResponseMessage(ctx, "success", "reminder", "get", http.StatusOK, res, metadata)
}
