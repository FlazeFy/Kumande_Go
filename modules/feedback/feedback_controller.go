package feedback

import (
	"kumande/models"
	"kumande/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	FeedbackService FeedbackService
}

func NewFeedbackController(feedbackService FeedbackService) *FeedbackController {
	return &FeedbackController{FeedbackService: feedbackService}
}

// Command
func (c *FeedbackController) CreateFeedback(ctx *gin.Context) {
	// Models
	var req models.Feedback

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		formattedErrors := utils.BuildValidationError(err)
		utils.BuildResponseMessage(ctx, "failed", "feedback", formattedErrors, http.StatusBadRequest, nil, nil)
		return
	}

	// Get User ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.BuildResponseMessage(ctx, "failed", "feedback", err.Error(), http.StatusBadRequest, nil, nil)
		return
	}

	// Service : Add Feedback
	err = c.FeedbackService.CreateFeedback(&req, *userID)
	if err != nil {
		utils.BuildErrorMessage(ctx, err.Error())
		return
	}

	utils.BuildResponseMessage(ctx, "success", "feedback", "post", http.StatusCreated, nil, nil)
}
