package hydration

import (
	"kumande/models"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HydrationController struct {
	HydrationService HydrationService
}

func NewHydrationController(hydrationService HydrationService) *HydrationController {
	return &HydrationController{HydrationService: hydrationService}
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
