package repositories

import (
	"context"
	"log"
	"subscription-service/database"
	"time"
)

func UpdateExpiredSubscriptions() error {
	query := `
		UPDATE subscriptions
		SET status = 'EXPIRED'
		WHERE end_date < $1 AND status != 'EXPIRED'
	`
	// Get the current date
	now := time.Now()

	// Execute the query
	_, err := database.DB.Exec(context.Background(), query, now)
	if err != nil {
		log.Println("Error updating expired subscriptions:", err)
		return err
	}

	log.Println("Expired subscriptions updated successfully at", now)
	return nil
}
