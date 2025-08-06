package factories

import (
	"kumande/configs"
	"kumande/models"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/crypto/bcrypt"
)

func UserFactory(username, email, telegramUserId, password *string, isValid bool) models.User {
	var finalUsername string
	if username != nil && *username != "" {
		finalUsername = *username
	} else {
		finalUsername = gofakeit.Username()
	}

	var finalEmail string
	if email != nil && *email != "" {
		finalEmail = *email
	} else {
		finalEmail = gofakeit.Email()
	}

	var pwd string
	if password != nil && *password != "" {
		pwd = *password
	} else {
		pwd = "nopass123"
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	var finalTelegramUserId *string
	if telegramUserId != nil && *telegramUserId != "" {
		finalTelegramUserId = telegramUserId
	} else {
		finalTelegramUserId = nil
	}

	now := time.Now()
	start := now.AddDate(-60, 0, 0)
	end := now.AddDate(-8, 0, 0)
	bornAt := gofakeit.DateRange(start, end)

	var bloodTypes *string
	if gofakeit.Bool() {
		bloodType := gofakeit.RandomString(configs.BloodTypes)
		bloodTypes = &bloodType
	} else {
		bloodTypes = nil
	}

	return models.User{
		Username:        finalUsername,
		Password:        string(hashedPass),
		TelegramUserId:  finalTelegramUserId,
		TelegramIsValid: isValid,
		ActivityFactor:  gofakeit.Float64Range(1.1, 2),
		Email:           finalEmail,
		Gender:          gofakeit.RandomString(configs.Genders),
		BornAt:          bornAt,
		BloodType:       bloodTypes,
		Currency:        gofakeit.RandomString(configs.Currencies),
	}
}
