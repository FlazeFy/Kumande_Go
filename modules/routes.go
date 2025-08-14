package modules

import (
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
	"kumande/modules/tag"
	userTrack "kumande/modules/user_track"
	userWeather "kumande/modules/user_weather"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB, redisClient *redis.Client,
	authController *auth.AuthController,
	feedbackController *feedback.FeedbackController,
	historyController *history.HistoryController,
	errorController *errors.ErrorController,
	dictionaryController *dictionary.DictionaryController,
	consumeController *consume.ConsumeController,
	userWeatherController *userWeather.UserWeatherController,
	reminderController *reminder.ReminderController,
	userTrackController *userTrack.UserTrackController,
	hydrationController *hydration.HydrationController,
	nutritionController *nutrition.NutritionController,
	allergicController *allergic.AllergicController,
	countCalorieController *countCalorie.CountCalorieController,
	bodyInfoController *bodyInfo.BodyInfoController,
	sleepController *sleep.SleepController,
	tagController *tag.TagController,
	budgetController *budget.BudgetController) {

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
	countCalorie.CountCalorieRouter(r, *countCalorieController, redisClient, db)
	bodyInfo.BodyInfoRouter(r, *bodyInfoController, redisClient, db)
	sleep.SleepRouter(r, *sleepController, redisClient, db)
	tag.TagRouter(r, *tagController, redisClient, db)
	budget.BudgetRouter(r, *budgetController, redisClient, db)
}
