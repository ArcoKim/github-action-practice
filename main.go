package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// HealthCheckResponse represents a response for the health check.
type HealthCheckResponse struct {
	Status string `json:"status"`
}

// ColorResponse represents a color response.
type ColorResponse struct {
	Color string `json:"color"`
}

// Colors is a slice of color names.
var Colors = []string{
	"Red", "Green", "Blue"
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/v1/color", colorHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// healthCheckHandler responds to health check requests.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{Status: "OK"}
	json.NewEncoder(w).Encode(response)
}

// colorHandler responds with a random color name.
func colorHandler(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(Colors))
	randomColor := Colors[randomIndex]
	response := ColorResponse{Color: randomColor}
	json.NewEncoder(w).Encode(response)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
