package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ServiceStatus struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func checkService(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "Down"
	}
	return "Up"
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")

	services := map[string]string{
		"jaypatel": "https://jaypatel.link",
		"magicdot": "https://magicdot.jaypatel.link",
		"dev":      "https://dev.jaypatel.link",
		"app":      "https://app.jaypatel.link",
		"res":      "https://res.jaypatel.link",
		"uptime":   "https://uptime.jaypatel.link",
	}

	var statuses []ServiceStatus
	for name, url := range services {
		status := checkService(url)
		uptime := "99.1"
		statuses = append(statuses, ServiceStatus{Name: name, Url: url, Status: status, Uptime: uptime})
	}
	json.NewEncoder(w).Encode(statuses)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", statusHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
