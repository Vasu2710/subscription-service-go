package routes

import (
	"subscription-service/controllers"
	"subscription-service/middlewares"

	"github.com/gorilla/mux"
)

func SetupRouter(secret_key string) *mux.Router {
	r := mux.NewRouter()
	secretKey := secret_key

	// User Routes
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	// Subscription Routes
	subscriptionRoutes := r.PathPrefix("/subscriptions").Subrouter()
	subscriptionRoutes.Use(middlewares.JWTMiddleware(secretKey)) // Apply middleware

	subscriptionRoutes.HandleFunc("", controllers.CreateSubscription).Methods("POST")
	subscriptionRoutes.HandleFunc("/{userId}", controllers.GetSubscription).Methods("GET")
	subscriptionRoutes.HandleFunc("/{userId}", controllers.UpdateSubscription).Methods("PUT")
	subscriptionRoutes.HandleFunc("/{userId}", controllers.CancelSubscription).Methods("DELETE")

	//Plan routes
	r.HandleFunc("/plans", controllers.GetPlans).Methods("GET")
	r.HandleFunc("/plans/{planId}", controllers.GetPlanDetails).Methods("GET")

	return r
}
