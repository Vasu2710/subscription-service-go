package cron

import (
	"log"
	"subscription-service/repositories"

	"github.com/robfig/cron"
)

func SetupCron() *cron.Cron {
	// Create a new cron instance
	c := cron.New()

	// Schedule the job to run every night at 12:00 AM
	err := c.AddFunc("0 0 * * *", func() {
		log.Println("Running cron job to check expired subscriptions")
		err := repositories.UpdateExpiredSubscriptions()
		if err != nil {
			log.Println("Error running subscription update job:", err)
		}
	})
	if err != nil {
		log.Fatal("Error adding cron job:", err)
	}

	// Start the cron scheduler
	c.Start()
	return c
}
