package countCalorie

import (
	"errors"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CountCalorieController struct {
	CountCalorieService CountCalorieService
}

func NewCountCalorieController(countCalorieService CountCalorieService) *CountCalorieController {
	return &CountCalorieController{CountCalorieService: countCalorieService}
}

// Query
func (c *CountCalorieController) GetLastCountCalorie(ctx *gin.Context) {
	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "count calorie", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get Count Calorie
	var data interface{}
	data, err = c.CountCalorieService.GetLastCountCalorie(*userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "count calorie", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	data = utils.StripFields(data, "created_by")
	utils.BuildResponseMessage(ctx, "success", "count calorie", "get", http.StatusOK, data, nil)
}

func (c *CountCalorieController) HardDeleteCountCalorieById(ctx *gin.Context) {
	// Param
	id := ctx.Param("id")

	// Parse Param UUID
	countCalorieID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "count calorie", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "count calorie", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	err = c.CountCalorieService.HardDeleteCountCalorieByID(countCalorieID, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "count calorie", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "count calorie", "hard delete", http.StatusOK, nil, nil)
}
