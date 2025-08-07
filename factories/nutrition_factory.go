package factories

import (
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func NutritionFactory() models.Nutrition {
	return models.Nutrition{
		CalorieMax: gofakeit.Number(10, 30) * 120,
		FatMax:     gofakeit.Float64Range(30.0, 120.0),
		ProteinMin: gofakeit.Float64Range(40.0, 150.0),
		CarbMax:    gofakeit.Float64Range(100.0, 300.0),
		SugarMax:   gofakeit.Float64Range(25.0, 100.0),
		SodiumMax:  gofakeit.Float64Range(1500.0, 4000.0),
		FiberMin:   gofakeit.Float64Range(15.0, 30.0),
	}
}
