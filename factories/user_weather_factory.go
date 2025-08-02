package factories

import (
	"kumande/configs"
	"kumande/models"
	"kumande/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func UserWeatherFactory() models.UserWeather {
	return models.UserWeather{
		WeatherTemp:      utils.GetRandWeatherTemp(-25.0, 45.0),
		WeatherHumid:     gofakeit.Number(10, 100),
		WeatherCity:      gofakeit.City(),
		WeatherCondition: gofakeit.RandomString(configs.WeatherConditions),
		WeatherHitFrom:   gofakeit.RandomString(configs.WeatherHitFroms),
	}
}
