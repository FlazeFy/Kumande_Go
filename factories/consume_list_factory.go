package factories

import (
	"encoding/json"
	"kumande/models"
	"kumande/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func ConsumeListFactory() models.ConsumeList {
	var consumeListTag []byte
	if gofakeit.Bool() {
		tagName := gofakeit.Word()
		slugName := utils.ConvertToSlug(tagName)
		checkpoints := []models.ConsumeTag{
			{TagName: tagName, SlugName: slugName},
			{TagName: tagName, SlugName: slugName},
			{TagName: tagName, SlugName: slugName},
		}
		jsonData, _ := json.Marshal(checkpoints)
		consumeListTag = jsonData
	} else {
		consumeListTag = nil
	}

	var consumeListDetail *string
	if gofakeit.Bool() {
		consumeDetailDum := gofakeit.Sentence(gofakeit.Number(3, 10))
		consumeListDetail = &consumeDetailDum
	} else {
		consumeListDetail = nil
	}

	return models.ConsumeList{
		ConsumeListName:   gofakeit.Sentence(gofakeit.Number(2, 4)),
		ConsumeListDetail: consumeListDetail,
		ConsumeListTag:    consumeListTag,
	}
}
