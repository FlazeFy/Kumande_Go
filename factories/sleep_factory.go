package factories

import (
	"kumande/models"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func SleepFactory() models.Sleep {
	baseDate := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())

	// Random sleep hour between 21:00 (9PM) and 02:00 (2AM next day)
	sleepHour := gofakeit.Number(21, 26)
	sleepTime := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), sleepHour%24, 0, 0, 0, time.UTC)

	// Woke At = 2 to 8 hours after sleep
	wokeTime := sleepTime.Add(time.Duration(gofakeit.Number(2, 8)) * time.Hour)

	var sleepNote *string
	if gofakeit.Bool() {
		sleepNoteDum := gofakeit.Sentence(gofakeit.Number(3, 10))
		sleepNote = &sleepNoteDum
	} else {
		sleepNote = nil
	}

	return models.Sleep{
		SleepNote:    sleepNote,
		SleepQuality: gofakeit.Number(1, 10),
		SleepAt:      sleepTime,
		WokeAt:       &wokeTime,
	}
}
