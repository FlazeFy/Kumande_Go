package consume

import (
	"errors"
	"kumande/configs"
	"kumande/modules/stats"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConsumeController struct {
	ConsumeService ConsumeService
	StatsService   stats.StatsService
}

func NewConsumeController(consumeService ConsumeService, statsService stats.StatsService) *ConsumeController {
	return &ConsumeController{
		ConsumeService: consumeService,
		StatsService:   statsService,
	}
}

func (c *ConsumeController) GetMostContextConsume(ctx *gin.Context) {
	// Param
	targetCol := ctx.Param("target_col")

	// Validator : Target Column Validator
	if !utils.Contains(configs.StatsConsumeField, targetCol) {
		utils.BuildResponseMessage(ctx, "failed", "consume", "target_col is not valid", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "consume", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	consume, err := c.StatsService.GetMostUsedContext("consumes", targetCol, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "consume", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "consume", "get", http.StatusOK, consume, nil)
}
