package feedback

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Feedback Interface
type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error
}

// Feedback Struct
type feedbackRepository struct {
	db *gorm.DB
}

// Feedback Constructor
func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error {
	// Default
	feedback.ID = uuid.New()
	feedback.CreatedAt = time.Now()
	feedback.CreatedBy = userID

	// Query
	return r.db.Create(feedback).Error
}
