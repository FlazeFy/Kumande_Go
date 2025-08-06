package factories

import (
	"encoding/json"
	"kumande/configs"
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func ConsumeFactory() models.Consume {
	consumeProvide := gofakeit.Company()
	if len(consumeProvide) > 36 {
		consumeProvide = consumeProvide[:36]
	}
	consumePrice := gofakeit.Number(2, 200) * 10000
	end := time.Now().AddDate(0, 0, -1)
	start := end.AddDate(0, 0, -60)
	consumeBuyAt := gofakeit.DateRange(start, end)

	var consumeTag []byte
	if gofakeit.Bool() {
		tagName := gofakeit.Word()
		slugName := utils.ConvertToSlug(tagName)
		checkpoints := []models.ConsumeTag{
			{TagName: tagName, SlugName: slugName},
			{TagName: tagName, SlugName: slugName},
			{TagName: tagName, SlugName: slugName},
		}
		jsonData, _ := json.Marshal(checkpoints)
		consumeTag = jsonData
	} else {
		consumeTag = nil
	}

	var consumeDetail *string
	if gofakeit.Bool() {
		consumeDetailDum := gofakeit.Sentence(gofakeit.Number(3, 10))
		consumeDetail = &consumeDetailDum
	} else {
		consumeDetail = nil
	}

	var consumeCal *int
	if gofakeit.Bool() {
		cal := gofakeit.Number(2, 50) * 10
		consumeCal = &cal
	} else {
		consumeCal = nil
	}

	return models.Consume{
		ConsumeName:    gofakeit.ProductName(),
		ConsumeDetail:  consumeDetail,
		ConsumeType:    gofakeit.RandomString(configs.ConsumeTypes),
		ConsumeFrom:    gofakeit.RandomString(configs.ConsumeFroms),
		ConsumePrice:   &consumePrice,
		ConsumeBuyAt:   &consumeBuyAt,
		ConsumeCal:     consumeCal,
		ConsumeQty:     gofakeit.Number(1, 3),
		ConsumeImage:   nil,
		ConsumeProvide: &consumeProvide,
		ConsumeTag:     consumeTag,
		IsFavorite:     gofakeit.Bool(),
	}
}
