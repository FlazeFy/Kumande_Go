package modules

import (
	"kumande/modules/admin"
	"kumande/modules/auth"
	"kumande/modules/budget"
	"kumande/modules/dictionary"
	"kumande/modules/errors"
	"kumande/modules/feedback"
	"kumande/modules/history"
	"kumande/modules/tag"
	"kumande/modules/user"
	"kumande/seeders"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	userRepo := user.NewUserRepository(db)
	feedbackRepo := feedback.NewFeedbackRepository(db)
	historyRepo := history.NewHistoryRepository(db)
	errorRepo := errors.NewErrorRepository(db)
	dictionaryRepo := dictionary.NewDictionaryRepository(db)
	budgetRepo := budget.NewBudgetRepository(db)
	tagRepo := tag.NewTagRepository(db)

	// Dependency Services
	adminService := admin.NewAdminService(adminRepo)
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)
	feedbackService := feedback.NewFeedbackService(feedbackRepo)
	historyService := history.NewHistoryService(historyRepo)
	errorService := errors.NewErrorService(errorRepo)
	dictionaryService := dictionary.NewDictionaryService(dictionaryRepo)

	// Dependency Controller
	authController := auth.NewAuthController(authService)
	feedbackController := feedback.NewFeedbackController(feedbackService)
	historyController := history.NewHistoryController(historyService)
	errorController := errors.NewErrorController(errorService)
	dictionaryController := dictionary.NewDictionaryController(dictionaryService)

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
	feedback.FeedbackRouter(r, *feedbackController, redisClient, db)
	history.HistoryRouter(r, *historyController, redisClient, db)
	errors.ErrorRouter(r, *errorController, redisClient, db)
	dictionary.DictionaryRouter(r, *dictionaryController, redisClient, db)

	// Task Scheduler
	SetUpScheduler(adminService)

	// Seeder & Factories
	seeders.SeedAdmins(adminRepo, 5)
	seeders.SeedUsers(userRepo, 20)
	seeders.SeedDictionaries(dictionaryRepo)
	seeders.SeedHistories(historyRepo, userRepo, 5)
	seeders.SeedBudget(budgetRepo, userRepo, 20)
	seeders.SeedTags(tagRepo, userRepo, 20)
}
