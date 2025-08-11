package modules

import (
	"kumande/modules/admin"
	"kumande/modules/allergic"
	"kumande/modules/auth"
	bodyInfo "kumande/modules/body_info"
	"kumande/modules/budget"
	"kumande/modules/consume"
	countCalorie "kumande/modules/count_calorie"
	"kumande/modules/dictionary"
	"kumande/modules/errors"
	"kumande/modules/feedback"
	"kumande/modules/history"
	"kumande/modules/hydration"
	"kumande/modules/nutrition"
	"kumande/modules/reminder"
	"kumande/modules/sleep"
	"kumande/modules/stats"
	"kumande/modules/tag"
	"kumande/modules/user"
	userTrack "kumande/modules/user_track"
	userWeather "kumande/modules/user_weather"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Cache
	statsCache := stats.NewStatsCache(redisClient)

	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	userRepo := user.NewUserRepository(db)
	feedbackRepo := feedback.NewFeedbackRepository(db)
	historyRepo := history.NewHistoryRepository(db)
	errorRepo := errors.NewErrorRepository(db)
	dictionaryRepo := dictionary.NewDictionaryRepository(db)
	budgetRepo := budget.NewBudgetRepository(db)
	tagRepo := tag.NewTagRepository(db)
	allergicRepo := allergic.NewAllergicRepository(db)
	bodyInfoRepo := bodyInfo.NewBodyInfoRepository(db)
	userTrackRepo := userTrack.NewUserTrackRepository(db)
	userWeatherRepo := userWeather.NewUserWeatherRepository(db)
	consumeRepo := consume.NewConsumeRepository(db)
	consumeListRepo := consume.NewConsumeListRepository(db)
	consumeListRelRepo := consume.NewConsumeListRelRepository(db)
	countCalorieRepo := countCalorie.NewCountCalorieRepository(db)
	reminderRepo := reminder.NewReminderRepository(db)
	reminderUsedRepo := reminder.NewReminderUsedRepository(db)
	statsRepo := stats.NewStatsRepository(db)
	sleepRepo := sleep.NewSleepRepository(db)
	hydrationRepo := hydration.NewHydrationRepository(db)
	consumeRateRepo := consume.NewConsumeRateRepository(db)
	nutritionRepo := nutrition.NewNutritionRepository(db)

	// Dependency Services
	adminService := admin.NewAdminService(adminRepo)
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)
	feedbackService := feedback.NewFeedbackService(feedbackRepo)
	historyService := history.NewHistoryService(historyRepo)
	errorService := errors.NewErrorService(errorRepo)
	dictionaryService := dictionary.NewDictionaryService(dictionaryRepo)
	userWeatherService := userWeather.NewUserWeatherService(userWeatherRepo)
	consumeService := consume.NewConsumeService(consumeRepo)
	statsService := stats.NewStatsService(statsRepo, redisClient, statsCache)
	reminderService := reminder.NewReminderService(reminderRepo)
	userTrackService := userTrack.NewUserTrackService(userTrackRepo)
	hydrationService := hydration.NewHydrationService(hydrationRepo)
	nutritionService := nutrition.NewNutritionService(nutritionRepo)
	allergicService := allergic.NewAllergicService(allergicRepo)
	countCalorieService := countCalorie.NewCountCalorieService(countCalorieRepo)
	bodyInfoService := bodyInfo.NewBodyInfoService(bodyInfoRepo)
	sleepService := sleep.NewSleepService(sleepRepo)
	tagService := tag.NewTagService(tagRepo)

	// Dependency Controller
	authController := auth.NewAuthController(authService)
	feedbackController := feedback.NewFeedbackController(feedbackService)
	historyController := history.NewHistoryController(historyService)
	errorController := errors.NewErrorController(errorService)
	dictionaryController := dictionary.NewDictionaryController(dictionaryService)
	consumeController := consume.NewConsumeController(consumeService, statsService)
	userWeatherController := userWeather.NewUserWeatherController(userWeatherService, statsService)
	reminderController := reminder.NewReminderController(reminderService, statsService)
	userTrackController := userTrack.NewUserTrackController(userTrackService, statsService)
	hydrationController := hydration.NewHydrationController(hydrationService)
	nutritionController := nutrition.NewNutritionController(nutritionService)
	allergicController := allergic.NewAllergicController(allergicService)
	countCalorieController := countCalorie.NewCountCalorieController(countCalorieService)
	bodyInfoController := bodyInfo.NewBodyInfoController(bodyInfoService)
	sleepController := sleep.NewSleepController(sleepService)
	tagController := tag.NewTagController(tagService)

	// Routes Endpoint
	SetUpRoutes(r, db, redisClient, authController, feedbackController, historyController, errorController, dictionaryController, consumeController, userWeatherController,
		reminderController, userTrackController, hydrationController, nutritionController, allergicController, countCalorieController, bodyInfoController, sleepController, tagController)

	// Task Scheduler
	SetUpScheduler(adminService, errorService, userWeatherService, userService, historyService)

	// Seeder & Factories
	SetUpSeeder(adminRepo, userRepo, dictionaryRepo, historyRepo, budgetRepo, tagRepo, errorRepo, allergicRepo, bodyInfoRepo, feedbackRepo, userTrackRepo, userWeatherRepo, consumeRepo, consumeListRepo, countCalorieRepo,
		consumeListRelRepo, reminderRepo, reminderUsedRepo, sleepRepo, hydrationRepo, consumeRateRepo, nutritionRepo)
}
