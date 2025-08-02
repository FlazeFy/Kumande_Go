package factories

import (
	"fmt"
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func BodyInfoFactory() models.BodyInfo {
	bloodPressure := fmt.Sprintf("%d/%d", gofakeit.Number(90, 190), gofakeit.Number(70, 150))

	return models.BodyInfo{
		BloodPressure: bloodPressure,
		BloodGlucose:  gofakeit.Number(70, 200),
		Gout:          gofakeit.Float32Range(2.0, 8.0),
		Cholesterol:   gofakeit.Number(130, 350),
	}
}
