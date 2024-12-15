package services

import (
	"errors"
	"fmt"
	"subscription-service/models"
	"subscription-service/repositories"
	"subscription-service/utils"
	"time"
)

func CreateSubscription(tokenStr string, secretKey string, sub models.Subscription) error {
	// Validate JWT token and get claims
	claims, err := utils.ValidateJWT(tokenStr, secretKey)
	fmt.Println(claims)
	if err != nil {
		return errors.New("unauthorized: invalid token")
	}

	// Extract userID from claims
	userIDFromToken, ok := claims["user_id"].(string)

	if !ok {
		return errors.New("unauthorized: invalid token claims")
	}

	// Ensure the userID in the token matches the userID in the subscription
	if sub.UserID != userIDFromToken {
		return errors.New("unauthorized: cannot create subscription for another user")
	}

	// Call repository to create the subscription
	return repositories.CreateSubscription(sub)
}

func GetSubscriptionByUserID(userID string) (models.Subscription, error) {
	return repositories.GetSubscriptionByUserID(userID)
}

func UpdateSubscription(userID string, newPlanID string) error {
	// Fetch the subscription ID using the user ID
	subscriptionID, err := repositories.GetSubscriptionIDByUserID(userID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return errors.New("no subscription found for the given user ID")
		}
		return err
	}

	// Calculate new fields
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)
	status := "ACTIVE"

	// Update the subscription
	err = repositories.UpdateSubscription(subscriptionID, newPlanID, status, startDate, endDate)
	if err != nil {
		return err
	}

	return nil
}

func CancelSubscription(userID string) error {
	return repositories.CancelSubscription(userID)
}

func ExpireSubscriptions() error {
	return repositories.ExpireSubscriptions()
}
