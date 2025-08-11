package sleep

import (
	"errors"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SleepController struct {
	SleepService SleepService
}

func NewSleepController(sleepService SleepService) *SleepController {
	return &SleepController{SleepService: sleepService}
}

// Command
func (c *SleepController) HardDeleteSleepById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	sleepID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "sleep", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "sleep", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Sleep By ID
	err = c.SleepService.HardDeleteSleepByID(sleepID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "sleep", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "sleep", "hard delete", http.StatusOK, nil, nil)
}
