package utils

import (
	"strings"
	"time"
)

func CalculateCalories(weight int, height int, age int, gender string, activityFactor float64) int {
	gender = strings.ToLower(gender)

	var bmr float64
	if gender == "male" {
		bmr = 10*float64(weight) + 6.25*float64(height) - 5*float64(age) + 5
	} else if gender == "female" {
		bmr = 10*float64(weight) + 6.25*float64(height) - 5*float64(age) - 161
	} else {
		return 0
	}

	return int(bmr * activityFactor)
}

func CalculateAge(bornDate time.Time) int {
	today := time.Now()
	age := today.Year() - bornDate.Year()

	if today.Month() < bornDate.Month() || (today.Month() == bornDate.Month() && today.Day() < bornDate.Day()) {
		age--
	}

	return age
}
