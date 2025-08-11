package bodyInfo

import (
	"errors"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BodyInfoController struct {
	BodyInfoService BodyInfoService
}

func NewBodyInfoController(bodyInfoService BodyInfoService) *BodyInfoController {
	return &BodyInfoController{BodyInfoService: bodyInfoService}
}

func (c *BodyInfoController) HardDeleteBodyInfoById(ctx *gin.Context) {
	// Param
	id := ctx.Param("id")

	// Parse Param UUID
	bodyInfoID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "body info", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "body info", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	err = c.BodyInfoService.HardDeleteBodyInfoByID(bodyInfoID, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "body info", "empty", http.StatusNotFound, nil, nil)
		default:
			utils.BuildErrorMessage(ctx, err.Error())
		}
		return
	}

	// Response
	utils.BuildResponseMessage(ctx, "success", "body info", "hard delete", http.StatusOK, nil, nil)
}
