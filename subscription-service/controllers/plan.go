package controllers

import (
	"encoding/json"
	"net/http"
	"subscription-service/services"

	"github.com/gorilla/mux"
)

func GetPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := services.GetPlans()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(plans)
}

func GetPlanDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planId := params["planId"]
	plan, err := services.GetPlanDetails(planId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(plan)
}
