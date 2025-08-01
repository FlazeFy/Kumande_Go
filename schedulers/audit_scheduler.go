package schedulers

import (
	"fmt"
	"kumande/modules/admin"
	"kumande/modules/errors"
	"kumande/utils"
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type AuditScheduler struct {
	ErrorService errors.ErrorService
	AdminService admin.AdminService
}

func NewAuditScheduler(
	errorService errors.ErrorService,
	adminService admin.AdminService,
) *AuditScheduler {
	return &AuditScheduler{
		ErrorService: errorService,
		AdminService: adminService,
	}
}

func (s *AuditScheduler) SchedulerAuditError() {
	// Service : Get All Admin Contact
	contact, err := s.AdminService.GetAllAdminContact()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Service : Get All Error Audit
	errorsList, err := s.ErrorService.SchedulerGetAllErrorAudit()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Send to Telegram
	datetime := time.Now()
	if len(contact) > 0 && len(errorsList) > 0 {
		filename := fmt.Sprintf("audit_error_%s.pdf", datetime)
		err = utils.GeneratePDFErrorAudit(errorsList, filename)
		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, dt := range contact {
			if dt.TelegramUserId != nil && dt.TelegramIsValid {
				bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
				if err != nil {
					log.Println("Failed to connect to Telegram bot")
					return
				}

				telegramID, err := strconv.ParseInt(*dt.TelegramUserId, 10, 64)
				if err != nil {
					log.Println("Invalid Telegram User Id")
					return
				}

				doc := tgbotapi.NewDocumentUpload(telegramID, filename)
				doc.ParseMode = "html"
				doc.Caption = fmt.Sprintf("[ADMIN] Hello %s, the system just run an audit error, with result of %d error found. Here's the document", dt.Username, len(errorsList))

				_, err = bot.Send(doc)
				if err != nil {
					log.Println(err.Error())
					return
				}
			}
		}

		// Cleanup
		os.Remove(filename)
	}
}
