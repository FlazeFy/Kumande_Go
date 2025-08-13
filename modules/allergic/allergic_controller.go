package allergic

import (
	"errors"
	"kumande/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AllergicController struct {
	AllergicService AllergicService
}

func NewAllergicController(allergicService AllergicService) *AllergicController {
	return &AllergicController{AllergicService: allergicService}
}

// Query
func (c *AllergicController) GetAllAllergic(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "allergic", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Allergic
	res, total, err := c.AllergicService.GetAllAllergic(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "allergic", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "allergic", "get", http.StatusOK, res, metadata)
}

// Command
func (c *AllergicController) HardDeleteAllergicById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	allergicID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "allergic", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "allergic", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Allergic By ID
	err = c.AllergicService.HardDeleteAllergicByID(allergicID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "allergic", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "allergic", "hard delete", http.StatusOK, nil, nil)
}
