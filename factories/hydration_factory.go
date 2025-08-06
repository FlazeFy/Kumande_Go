package factories

import (
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func HydrationFactory() models.Hydration {
	return models.Hydration{
		VolumeML: gofakeit.Number(5, 9) * 50,
	}
}
