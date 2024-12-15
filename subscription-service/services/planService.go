package services

import (
	"context"
	"encoding/json"
	"fmt"
	"subscription-service/database"
	"subscription-service/models"
	"time"
)

func GetPlans() ([]models.Plan, error) {
	// Redis cache key
	cacheKey := "plans:all"

	// Check Redis cache
	cachedPlans, err := database.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil && cachedPlans != "" {
		// Cache hit: Unmarshal the cached data
		var plans []models.Plan
		err = json.Unmarshal([]byte(cachedPlans), &plans)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached plans: %v", err)
		}
		fmt.Println("Serving from cache")
		return plans, nil
	}

	// Cache miss: Query the database
	query := `SELECT id, name, price, features, duration FROM plans`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse rows into plans
	var plans []models.Plan
	for rows.Next() {
		var plan models.Plan
		err := rows.Scan(&plan.ID, &plan.Name, &plan.Price, &plan.Features, &plan.Duration)
		if err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}

	// Check for row iteration errors
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	// Serialize the plans and store them in Redis
	plansJSON, err := json.Marshal(plans)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal plans: %v", err)
	}

	err = database.RedisClient.Set(context.Background(), cacheKey, plansJSON, time.Hour).Err() // Cache for 1 hour
	if err != nil {
		fmt.Printf("Failed to cache plans: %v\n", err)
	}

	fmt.Println("Serving from database")
	return plans, nil
}

func GetPlanDetails(planId string) (models.Plan, error) {
	// Logic to retrieve plan details from database
	return models.Plan{}, nil
}
