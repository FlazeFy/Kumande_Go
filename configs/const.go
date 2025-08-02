package configs

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
