package modules

import (
	"kumande/modules/admin"
	"kumande/modules/errors"
	"kumande/modules/user"
	userWeather "kumande/modules/user_weather"
	"kumande/schedulers"
	"time"

	"github.com/robfig/cron"
)

func SetUpScheduler(adminService admin.AdminService, errorService errors.ErrorService, userWeatherService userWeather.UserWeatherService, userService user.UserService) {
	// Initialize Scheduler
	houseKeepingScheduler := schedulers.NewHouseKeepingScheduler(adminService)
	auditScheduler := schedulers.NewAuditScheduler(errorService, adminService)
	weatherScheduler := schedulers.NewWeatherScheduler(adminService, userService, userWeatherService)

	// Init Scheduler
	c := cron.New()
	Scheduler(c, houseKeepingScheduler, auditScheduler, weatherScheduler)
	c.Start()
	defer c.Stop()
}

func Scheduler(c *cron.Cron, houseKeepingScheduler *schedulers.HouseKeepingScheduler, auditScheduler *schedulers.AuditScheduler, weatherScheduler *schedulers.WeatherScheduler) {
	// For Production
	// Audit Scheduler
	c.AddFunc("0 1 * * 1", auditScheduler.SchedulerAuditError)

	// House Keeping Scheduler
	c.AddFunc("0 5 2 * *", houseKeepingScheduler.SchedulerMonthlyLog)

	// Weather Scheduler
	c.AddFunc("10 0 * * *", weatherScheduler.SchedulerWeatherRoutineFetch)

	// For Development
	go func() {
		time.Sleep(5 * time.Second)

		// Audit Scheduler
		auditScheduler.SchedulerAuditError()

		// House Keeping Scheduler
		houseKeepingScheduler.SchedulerMonthlyLog()

		// Weather Scheduler
		weatherScheduler.SchedulerWeatherRoutineFetch()
	}()
}
