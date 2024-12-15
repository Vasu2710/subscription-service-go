package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"subscription-service/models"
	"subscription-service/services"

	"github.com/gorilla/mux"
)

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header
	authHeader := r.Header.Get("Authorization")
	fmt.Println(authHeader, "rfrfrf")
	if authHeader == "" {
		http.Error(w, "missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Extract the Bearer token from the header
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "invalid Authorization header format", http.StatusUnauthorized)
		return
	}
	token := parts[1]
	fmt.Println(token, "cdcdcdcd")
	// Parse the subscription details from the request body
	var sub models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Secret key for JWT (should ideally come from environment variables)
	secretKey := os.Getenv("JWT_SECRET")

	// Call the service layer to create the subscription
	if err := services.CreateSubscription(token, secretKey, sub); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Subscription created successfully"))
}
func GetSubscription(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]
	sub, err := services.GetSubscriptionByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(sub)
}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the URL
	vars := mux.Vars(r)
	userID := vars["userId"]
	fmt.Print(userID)
	// Parse the request body for the new plan ID
	var requestBody struct {
		PlanID string `json:"planId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the plan ID
	if requestBody.PlanID == "" {
		http.Error(w, "Plan ID is required", http.StatusBadRequest)
		return
	}

	// Call the service layer to update the subscription
	err := services.UpdateSubscription(userID, requestBody.PlanID)
	if err != nil {
		if err.Error() == "no subscription found for the given user ID" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update subscription", http.StatusInternalServerError)
		}
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Subscription updated successfully"))
}
func CancelSubscription(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]
	if err := services.CancelSubscription(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ExpireSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	if err := services.ExpireSubscriptions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
