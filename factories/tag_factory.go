package factories

import (
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func TagFactory() models.Tag {
	return models.Tag{
		TagName: gofakeit.Word(),
	}
}
