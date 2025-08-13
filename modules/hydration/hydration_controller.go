package hydration

import (
	"errors"
	"kumande/models"
	"kumande/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HydrationController struct {
	HydrationService HydrationService
}

func NewHydrationController(hydrationService HydrationService) *HydrationController {
	return &HydrationController{HydrationService: hydrationService}
}

// Query
func (c *HydrationController) GetHydrationByDate(ctx *gin.Context) {
	// Param
	date := ctx.Param("date")

	// Validate Date
	if !utils.ValidateDateFormat(date, "02-01-2006") {
		utils.BuildResponseMessage(ctx, "failed", "hydration", "invalid date format", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "hydration", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get Hydration Last History
	// Note : date format DD-MM-YYYY
	var data interface{}
	data, err = c.HydrationService.GetHydrationByDate(*userID, date)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "hydration", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	data = utils.StripFields(data, "created_by")
	utils.BuildResponseMessage(ctx, "success", "hydration", "get", http.StatusOK, data, nil)
}

func (c *HydrationController) GetAllHydration(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "hydration", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Hydration
	res, total, err := c.HydrationService.GetAllHydration(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "hydration", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "hydration", "get", http.StatusOK, res, metadata)
}

// Command
func (c *HydrationController) PostCreateHydration(ctx *gin.Context) {
	// Models
	var req models.Hydration

	// Validate JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		formattedErrors := utils.BuildValidationError(err)
		utils.BuildResponseMessage(ctx, "failed", "hydration", formattedErrors, http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "hydration", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Create Hydration
	hydration := models.Hydration{
		VolumeML: req.VolumeML,
	}
	err = c.HydrationService.CreateHydration(hydration, *userID)
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "hydration", "post", http.StatusCreated, nil, nil)
}

func (c *HydrationController) HardDeleteHydrationById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	hydrationID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "hydration", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "hydration", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Hydration By ID
	err = c.HydrationService.HardDeleteHydrationByID(hydrationID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "hydration", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "hydration", "hard delete", http.StatusOK, nil, nil)
}
