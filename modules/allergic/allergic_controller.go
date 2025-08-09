package allergic

import (
	"errors"
	"kumande/utils"
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
