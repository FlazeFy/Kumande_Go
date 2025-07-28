package modules

import (
	"kumande/modules/admin"
	"kumande/modules/auth"
	"kumande/modules/feedback"
	"kumande/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	userRepo := user.NewUserRepository(db)
	feedbackRepo := feedback.NewFeedbackRepository(db)

	// Dependency Services
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)
	feedbackService := feedback.NewFeedbackService(feedbackRepo)

	// Dependency Controller
	authController := auth.NewAuthController(authService)
	feedbackController := feedback.NewFeedbackController(feedbackService)

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
	feedback.FeedbackRouter(r, *feedbackController, redisClient, db)
}
