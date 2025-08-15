package budget

import (
	"errors"
	"kumande/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BudgetController struct {
	BudgetService BudgetService
}

func NewBudgetController(budgetService BudgetService) *BudgetController {
	return &BudgetController{BudgetService: budgetService}
}

// Queries
func (c *BudgetController) GetAllBudget(ctx *gin.Context) {
	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "consume", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Budget
	budget, total, err := c.BudgetService.GetAllBudget(pagination, *userID)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "budget", "get", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "budget", "get", http.StatusOK, budget, metadata)
}

func (c *BudgetController) GetBudgetByYear(ctx *gin.Context) {
	// Param
	year := ctx.Param("year")

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "consume", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get Budget By Year
	budget, err := c.BudgetService.GetBudgetByYear(year, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "budget", "get", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	utils.BuildResponseMessage(ctx, "success", "budget", "get", http.StatusOK, budget, nil)
}
