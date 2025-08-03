package schedulers

import (
	"fmt"
	"kumande/modules/admin"
	"kumande/modules/history"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CleanScheduler struct {
	AdminService   admin.AdminService
	HistoryService history.HistoryService
}

func NewCleanScheduler(
	adminService admin.AdminService,
	historyService history.HistoryService,
) *CleanScheduler {
	return &CleanScheduler{
		AdminService:   adminService,
		HistoryService: historyService,
	}
}

func (s *CleanScheduler) SchedulerCleanHistory() {
	days := 30

	// Service : Get All Admin Contact
	contact, err := s.AdminService.GetAllAdminContact()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Service : Delete History For Last N Days
	total, err := s.HistoryService.DeleteHistoryForLastNDays(days)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Send to Telegram
	if len(contact) > 0 {
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

				msgText := fmt.Sprintf("[ADMIN] Hello %s, the system just run a clean history, with result of %d history executed", dt.Username, total)
				msg := tgbotapi.NewMessage(telegramID, msgText)

				_, err = bot.Send(msg)
				if err != nil {
					log.Println("Failed to send message to Telegram")
					return
				}
			}
		}
	}
}
