package factories

import (
	"kumande/configs"
	"kumande/models"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func BudgetFactory() models.Budget {
	var overAt *time.Time
	if gofakeit.Bool() {
		t := gofakeit.DateRange(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC), time.Now())
		overAt = &t
	}

	return models.Budget{
		BudgetTotal: 100000 * gofakeit.Number(20, 75),
		BudgetMonth: gofakeit.RandomString(configs.Months),
		BudgetYear:  gofakeit.Number(2012, 2024),
		OverAt:      overAt,
	}
}
