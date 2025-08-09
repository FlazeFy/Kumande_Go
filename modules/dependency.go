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
	"kumande/seeders"

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

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
	feedback.FeedbackRouter(r, *feedbackController, redisClient, db)
	history.HistoryRouter(r, *historyController, redisClient, db)
	errors.ErrorRouter(r, *errorController, redisClient, db)
	dictionary.DictionaryRouter(r, *dictionaryController, redisClient, db)
	consume.ConsumeRouter(r, *consumeController, redisClient, db)
	userWeather.UserWeatherRouter(r, *userWeatherController, redisClient, db)
	reminder.ReminderRouter(r, *reminderController, redisClient, db)
	userTrack.UserTrackRouter(r, *userTrackController, redisClient, db)
	hydration.HydrationRouter(r, *hydrationController, redisClient, db)
	nutrition.NutritionRouter(r, *nutritionController, redisClient, db)
	allergic.AllergicRouter(r, *allergicController, redisClient, db)

	// Task Scheduler
	SetUpScheduler(adminService, errorService, userWeatherService, userService, historyService)

	// Seeder & Factories
	seeders.SeedAdmins(adminRepo, 5)
	seeders.SeedUsers(userRepo, 20)
	seeders.SeedDictionaries(dictionaryRepo)
	seeders.SeedHistories(historyRepo, userRepo, 5)
	seeders.SeedBudget(budgetRepo, userRepo, 20)
	seeders.SeedTags(tagRepo, userRepo, 20)
	seeders.SeedErrors(errorRepo, 20)
	seeders.SeedAllergic(allergicRepo, userRepo, 20)
	seeders.SeedBodyInfo(bodyInfoRepo, userRepo, 60)
	seeders.SeedFeedbacks(feedbackRepo, userRepo, 20)
	seeders.SeedUserTracks(userTrackRepo, userRepo, 60)
	seeders.SeedUserWeathers(userWeatherRepo, userRepo, 30)
	seeders.SeedConsume(consumeRepo, userRepo, 100)
	seeders.SeedConsumeList(consumeListRepo, userRepo, 50)
	seeders.SeedCountCalorie(countCalorieRepo, userRepo, 60)
	seeders.SeedConsumeListRelations(consumeListRelRepo, userRepo, consumeRepo, consumeListRepo, 10)
	seeders.SeedReminder(reminderRepo, userRepo, 20)
	seeders.SeedReminderUsed(reminderUsedRepo, userRepo, reminderRepo, 20)
	seeders.SeedSleeps(sleepRepo, userRepo, 60)
	seeders.SeedHydrations(hydrationRepo, userRepo, 120)
	seeders.SeedConsumeRates(consumeRateRepo, userRepo, consumeRepo, 20)
	seeders.SeedNutritions(nutritionRepo, userRepo, 60)
}
