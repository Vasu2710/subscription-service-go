package models

type Plan struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Features string  `json:"features"`
	Duration int     `json:"duration"`
}