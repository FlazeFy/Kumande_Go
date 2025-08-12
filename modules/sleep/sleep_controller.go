package sleep

import (
	"errors"
	"kumande/utils"
	"math"
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

// Query
func (c *SleepController) GetAllSleep(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "sleep", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Sleep
	res, total, err := c.SleepService.GetAllSleep(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "sleep", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "sleep", "get", http.StatusOK, res, metadata)
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
