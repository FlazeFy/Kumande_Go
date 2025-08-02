package modules

import (
	"kumande/modules/admin"
	"kumande/modules/errors"
	"kumande/schedulers"
	"time"

	"github.com/robfig/cron"
)

func SetUpScheduler(adminService admin.AdminService, errorService errors.ErrorService) {

	// Initialize Scheduler
	houseKeepingScheduler := schedulers.NewHouseKeepingScheduler(adminService)
	auditScheduler := schedulers.NewAuditScheduler(errorService, adminService)

	// Init Scheduler
	c := cron.New()
	Scheduler(c, houseKeepingScheduler, auditScheduler)
	c.Start()
	defer c.Stop()
}

func Scheduler(c *cron.Cron, houseKeepingScheduler *schedulers.HouseKeepingScheduler, auditScheduler *schedulers.AuditScheduler) {
	// For Production
	// Clean Scheduler
	// c.AddFunc("0 1 * * 1", auditScheduler.SchedulerAuditError)
	// c.AddFunc("0 5 2 * *", houseKeepingScheduler.SchedulerMonthlyLog)

	// For Development
	go func() {
		time.Sleep(5 * time.Second)

		// Audit Scheduler
		auditScheduler.SchedulerAuditError()

		// House Keeping Scheduler
		houseKeepingScheduler.SchedulerMonthlyLog()
	}()
}
