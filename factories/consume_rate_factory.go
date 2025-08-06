package factories

import (
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func ConsumeRateFactory(consumeID uuid.UUID) models.ConsumeRate {
	var consumeComment *string
	if gofakeit.Bool() {
		consumeCommentDum := gofakeit.Sentence(gofakeit.Number(3, 10))
		consumeComment = &consumeCommentDum
	} else {
		consumeComment = nil
	}

	return models.ConsumeRate{
		ConsumeComment: consumeComment,
		ConsumeRate:    gofakeit.Number(1, 10),
		ConsumeId:      consumeID,
	}
}
