package userTrack

import (
	"errors"
	"kumande/configs"
	"kumande/modules/stats"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserTrackController struct {
	UserTrackService UserTrackService
	StatsService     stats.StatsService
}

func NewUserTrackController(trackService UserTrackService, statsService stats.StatsService) *UserTrackController {
	return &UserTrackController{
		UserTrackService: trackService,
		StatsService:     statsService,
	}
}

// Query
func (c *UserTrackController) GetMostContextUserTrack(ctx *gin.Context) {
	// Param
	targetCol := ctx.Param("target_col")

	// Validator : Target Column Validator
	if !utils.Contains(configs.StatsUserTrackField, targetCol) {
		utils.BuildResponseMessage(ctx, "failed", "user track", "target_col is not valid", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "user track", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service: Get Most Context
	track, err := c.StatsService.GetMostUsedContext("user_tracks", targetCol, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "user track", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "user track", "get", http.StatusOK, track, nil)
}
