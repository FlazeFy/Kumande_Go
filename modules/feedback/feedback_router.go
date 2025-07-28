package feedback

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func FeedbackRouter(r *gin.Engine, feedbackController FeedbackController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Public Routes
		feedback := api.Group("/feedbacks")
		{
			feedback.POST("/", feedbackController.CreateFeedback)
		}
	}
}
