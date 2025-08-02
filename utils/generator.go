package utils

import (
	"math"

	"github.com/brianvoe/gofakeit/v6"
)

func GetRandWeatherTemp(min, max float64) float64 {
	raw := gofakeit.Float64Range(min, max)
	temp := math.Round(raw*100) / 100

	return temp
}
