package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uptime/models"
)

func GetStatusPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Print("insied the handler")
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Missing username parameter", http.StatusBadRequest)
			return
		}

		// statuses, err := services.GetServiceStatuses(db)
		// if err != nil {
		// 	http.Error(w, "Failed to fetch service statuses", http.StatusInternalServerError)
		// 	return
		// }
		lastDownTime, _ := time.Parse(time.RFC3339, "2025-01-10T22:27:46.640373Z")
		statuses := []models.ServiceStatus{
			{
				Name:            "test",
				Url:             "https://test.link",
				Status:          "Up",
				Uptime:          "6 days, 0 hours, 51 minutes",
				Uptime_percent:  "20",
				Uptime_duration: 521490225893021,
				// Last_down:       "2025-01-10T22:27:46.640373Z",
				Last_down: lastDownTime,
			},
		}

		json.NewEncoder(w).Encode(statuses)
	}
}
