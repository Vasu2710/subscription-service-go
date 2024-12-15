package repositories

import (
	"context"
	"fmt"
	"subscription-service/database"
	"subscription-service/models"
	"time"

	"github.com/google/uuid"
)

func CreateSubscription(sub models.Subscription) error {
	sub.ID = uuid.New().String()
	sub.StartDate = time.Now()
	sub.EndDate = sub.StartDate.AddDate(0, 0, 30)
	sub.Status = "ACTIVE"

	query := `INSERT INTO subscriptions (id, user_id, plan_id, status, start_date, end_date) 
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := database.DB.Exec(context.Background(), query, sub.ID, sub.UserID, sub.PlanID, sub.Status, sub.StartDate, sub.EndDate)
	return err
}

func GetSubscriptionByUserID(userID string) (models.Subscription, error) {
	var sub models.Subscription
	query := `SELECT id, user_id, plan_id, status, start_date, end_date FROM subscriptions WHERE user_id = $1`
	err := database.DB.QueryRow(context.Background(), query, userID).Scan(&sub.ID, &sub.UserID, &sub.PlanID, &sub.Status, &sub.StartDate, &sub.EndDate)
	return sub, err
}
func GetSubscriptionIDByUserID(userID string) (uuid.UUID, error) {
	// Query to fetch the subscription ID by user ID
	var subscriptionID uuid.UUID
	query := `SELECT id FROM subscriptions WHERE user_id = $1`
	fmt.Println("Executing query:", query, "with userID:", userID)
	userId, _ := uuid.Parse(userID)
	fmt.Print(userID)
	fmt.Print(userId)
	err := database.DB.QueryRow(context.Background(), query, userId).Scan(&subscriptionID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			fmt.Println("No subscription found for userID:", userID)
			return subscriptionID, err
		}
		fmt.Println("Error while querying subscription ID:", err)
		return subscriptionID, err
	}

	fmt.Println("Fetched subscription ID:", subscriptionID, "for userID:", userID)
	return subscriptionID, nil
}

func UpdateSubscription(subscriptionID uuid.UUID, planID, status string, startDate, endDate time.Time) error {
	// Query to update the subscription
	query := `UPDATE subscriptions SET plan_id = $1, status = $2, start_date = $3, end_date = $4 WHERE id = $5`
	_, err := database.DB.Exec(context.Background(), query, planID, status, startDate, endDate, subscriptionID)
	return err
}
func CancelSubscription(userID string) error {
	query := `UPDATE subscriptions SET status = 'CANCELLED' WHERE user_id = $1`
	_, err := database.DB.Exec(context.Background(), query, userID)
	return err
}

func ExpireSubscriptions() error {
	query := `UPDATE subscriptions SET status = 'EXPIRED' WHERE status = 'ACTIVE' AND end_date < $1`
	_, err := database.DB.Exec(context.Background(), query, time.Now())
	return err
}
