package tag

import (
	"errors"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagController struct {
	TagService TagService
}

func NewTagController(tagService TagService) *TagController {
	return &TagController{TagService: tagService}
}

// Command
func (c *TagController) HardDeleteTagById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Parse Param UUID
	tagID, err := uuid.Parse(id)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "tag", "invalid id", http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "tag", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Hard Delete Tag By ID
	err = c.TagService.HardDeleteTagByID(tagID, *userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.BuildResponseMessage(ctx, "failed", "tag", "empty", http.StatusNotFound, nil, nil)
		return
	}
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "tag", "hard delete", http.StatusOK, nil, nil)
}
