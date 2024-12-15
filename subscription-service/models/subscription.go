package models

import (
	"time"
)

type SubscriptionStatus string

const (
	StatusActive    SubscriptionStatus = "ACTIVE"
	StatusInactive  SubscriptionStatus = "INACTIVE"
	StatusCancelled SubscriptionStatus = "CANCELLED"
	StatusExpired   SubscriptionStatus = "EXPIRED"
)

type Subscription struct {
	ID        string             `json:"id"`
	UserID    string             `json:"userId"`
	PlanID    string             `json:"planId"`
	Status    SubscriptionStatus `json:"status"`
	StartDate time.Time          `json:"startDate"`
	EndDate   time.Time          `json:"endDate"`
}