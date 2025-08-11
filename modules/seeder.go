package modules

import (
	"kumande/modules/admin"
	"kumande/modules/allergic"
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
	"kumande/modules/user"
	userTrack "kumande/modules/user_track"
	userWeather "kumande/modules/user_weather"
	"kumande/seeders"
)

func SetUpSeeder(
	adminRepo admin.AdminRepository, userRepo user.UserRepository,
	dictionaryRepo dictionary.DictionaryRepository, historyRepo history.HistoryRepository,
	budgetRepo budget.BudgetRepository, tagRepo tag.TagRepository, errorRepo errors.ErrorRepository,
	allergicRepo allergic.AllergicRepository, bodyInfoRepo bodyInfo.BodyInfoRepository,
	feedbackRepo feedback.FeedbackRepository, userTrackRepo userTrack.UserTrackRepository,
	userWeatherRepo userWeather.UserWeatherRepository, consumeRepo consume.ConsumeRepository,
	consumeListRepo consume.ConsumeListRepository, countCalorieRepo countCalorie.CountCalorieRepository,
	consumeListRelRepo consume.ConsumeListRelRepository, reminderRepo reminder.ReminderRepository,
	reminderUsedRepo reminder.ReminderUsedRepository, sleepRepo sleep.SleepRepository,
	hydrationRepo hydration.HydrationRepository, consumeRateRepo consume.ConsumeRateRepository,
	nutritionRepo nutrition.NutritionRepository,
) {
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
