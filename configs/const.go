package configs

import "time"

var ResponseMessages = map[string]string{
	"post":        "created",
	"put":         "updated",
	"hard delete": "permanentally deleted",
	"soft delete": "deleted",
	"recover":     "recovered",
	"get":         "fetched",
	"login":       "login",
	"sign out":    "signed out",
	"empty":       "not found",
}
var Currencies = []string{"IDR", "USD", "EUR", "JPY", "GBP", "CNY", "CAD", "CHF", "AUD", "HKD", "SGD"}
var Genders = []string{"male", "female"}
var Months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var TrackSources = []string{"Web", "Mobile", "Telegram Bot", "Line Bot"}
var WeatherHitFroms = []string{"Task Schedule", "Manual"}
var WeatherConditions = []string{"Thunderstorm", "Drizzle", "Rain", "Snow", "Mist", "Smoke", "Haze", "Dust", "Fog", "Sand", "Ash", "Squall", "Tornado", "Clear", "Clouds"}
var ConsumeTypes = []string{"Food", "Drink", "Snack"}
var ConsumeFroms = []string{"GoFood", "GrabFood", "ShopeeFood", "Dine-In", "Take Away", "Cooking"}
var ReminderTypes = []string{"Every Day", "Every Month", "Every Year"}
var ReminderAttachmentTypes = []string{"Image", "URL", "Location"}
var StatsConsumeField = []string{"consume_type", "consume_from", "consume_provide"}
var StatsWeatherField = []string{"weather_condition", "weather_city"}
var BloodTypes = []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}

// Rules
var RedisTime = 10 * time.Minute
