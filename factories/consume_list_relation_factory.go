package factories

import (
	"kumande/models"

	"github.com/google/uuid"
)

func ConsumeListRelationFactory(consumeListID, consumeID uuid.UUID) models.ConsumeListRelation {
	return models.ConsumeListRelation{
		ConsumeListId: consumeListID,
		ConsumeId:     consumeID,
	}
}
