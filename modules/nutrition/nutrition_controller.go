package nutrition

import (
	"errors"
	"kumande/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NutritionController struct {
	NutritionService NutritionService
}

func NewNutritionController(nutritionService NutritionService) *NutritionController {
	return &NutritionController{NutritionService: nutritionService}
}

// Query
func (c *NutritionController) GetAllNutrition(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "nutrition", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Nutrition
	res, total, err := c.NutritionService.GetAllNutrition(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "nutrition", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "nutrition", "get", http.StatusOK, res, metadata)
}

// Command
func (c *NutritionController) HardDeleteNutritionById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	nutritionID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "nutrition", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "nutrition", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Nutrition By ID
	err = c.NutritionService.HardDeleteNutritionByID(nutritionID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "nutrition", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "nutrition", "hard delete", http.StatusOK, nil, nil)
}
