package factories

import (
	"encoding/json"
	"fmt"
	"kumande/configs"
	"kumande/models"

	"github.com/brianvoe/gofakeit/v6"
)

func ReminderFactory() models.Reminder {
	// Reminder Type
	reminderType := gofakeit.RandomString(configs.ReminderTypes)

	// Reminder Context
	var reminderContext []byte
	var contexts []models.ReminderContext
	totalReminderContext := gofakeit.Number(1, 4)

	switch reminderType {
	case "Every Day":
		for i := 0; i < totalReminderContext; i++ {
			hour := gofakeit.Number(0, 23)
			minute := []int{0, 30}[gofakeit.Number(0, 1)]
			time := fmt.Sprintf("%02d:%02d", hour, minute)
			contexts = append(contexts, models.ReminderContext{Time: time})
		}

	case "Every Month":
		for i := 0; i < totalReminderContext; i++ {
			day := gofakeit.Number(1, 30)
			time := fmt.Sprintf("Day %d", day)
			contexts = append(contexts, models.ReminderContext{Time: time})
		}

	case "Every Year":
		months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
		for i := 0; i < totalReminderContext; i++ {
			day := gofakeit.Number(1, 31)
			month := gofakeit.RandomString(months)
			time := fmt.Sprintf("%d %s", day, month)
			contexts = append(contexts, models.ReminderContext{Time: time})
		}
	}

	jsonData, _ := json.Marshal(contexts)
	reminderContext = jsonData

	// Reminder Attachment
	var attachments []models.ReminderAttachment
	totalReminderAttachment := gofakeit.Number(1, 2)
	attachmentType := gofakeit.RandomString(configs.ReminderAttachmentTypes)
	switch attachmentType {
	case "Image", "URL":
		for i := 0; i < totalReminderAttachment; i++ {
			attachments = append(attachments, models.ReminderAttachment{
				AttachmentType:    attachmentType,
				AttachmentName:    gofakeit.LoremIpsumSentence(2),
				AttachmentContext: fmt.Sprintf(`"%s"`, gofakeit.URL()),
			})
		}

	case "Location":
		for i := 0; i < totalReminderAttachment; i++ {
			latitude := gofakeit.Latitude()
			longitude := gofakeit.Longitude()
			coordinate := fmt.Sprintf("%.6f,%.6f", latitude, longitude)

			attachments = append(attachments, models.ReminderAttachment{
				AttachmentType:    attachmentType,
				AttachmentName:    gofakeit.LoremIpsumSentence(2),
				AttachmentContext: fmt.Sprintf(`"%s"`, coordinate),
			})
		}

	}

	jsonAttachmentData, _ := json.Marshal(attachments)
	reminderAttachment := jsonAttachmentData

	return models.Reminder{
		ReminderName:       gofakeit.LoremIpsumSentence(1),
		ReminderType:       reminderType,
		ReminderContext:    reminderContext,
		ReminderBody:       gofakeit.LoremIpsumSentence(3),
		ReminderAttachment: reminderAttachment,
	}
}
