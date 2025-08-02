package factories

import (
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func AllergicFactory() models.Allergic {
	randIntAllergicContext := gofakeit.Number(7, 15)
	allergicDesc := gofakeit.LoremIpsumSentence(randIntAllergicContext)

	return models.Allergic{
		AllergicContext: gofakeit.Word(),
		AllergicDesc:    &allergicDesc,
	}
}
