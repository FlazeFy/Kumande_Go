package factories

import (
	"kumande/models"
	"kumande/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func CountCalorieFactory(age int, gender string, activityFactor float64) models.CountCalorie {
	weight := gofakeit.Number(40, 90)
	height := gofakeit.Number(155, 200)
	result := utils.CalculateCalories(weight, height, age, gender, activityFactor)

	return models.CountCalorie{
		Weight: weight,
		Height: height,
		Result: int(result),
	}
}
