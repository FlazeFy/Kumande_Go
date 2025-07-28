package feedback

import (
	"kumande/models"

	"github.com/google/uuid"
)

// Feedback Interface
type FeedbackService interface {
	CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error
}

// Feedback Struct
type feedbackService struct {
	feedbackRepo FeedbackRepository
}

// Feedback Constructor
func NewFeedbackService(feedbackRepo FeedbackRepository) FeedbackService {
	return &feedbackService{
		feedbackRepo: feedbackRepo,
	}
}

func (r *feedbackService) CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error {
	return r.feedbackRepo.CreateFeedback(feedback, userID)
}
