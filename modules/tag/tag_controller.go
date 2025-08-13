package tag

import (
	"errors"
	"kumande/utils"
	"math"
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

// Query
func (c *TagController) GetAllTag(ctx *gin.Context) {
	var res interface{}

	// Pagination
	pagination := utils.GetPagination(ctx)

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "tag", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Get All Tag
	res, total, err := c.TagService.GetAllTag(pagination, *userID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			utils.BuildResponseMessage(ctx, "failed", "tag", "empty", http.StatusNotFound, nil, nil)
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
	utils.BuildResponseMessage(ctx, "success", "tag", "get", http.StatusOK, res, metadata)
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
