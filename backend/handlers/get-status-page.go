package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"uptime/models"
	"uptime/repository"
)

func GetStatusPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Print("inside the handler")
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Missing username parameter", http.StatusBadRequest)
			return
		}
		log.Print(username)

		project, err := repository.GetProjects(db, username)
		if err != nil {
			log.Print(err)
		}
		log.Print("test")
		log.Print(project)
		lastDownTime, _ := time.Parse(time.RFC3339, "2024-01-10T22:27:46.640373Z")
		statuses := []models.ServiceStatus{
			{
				Name:            project,
				Url:             project,
				Status:          "Up",
				Uptime:          "366 days, 0 hours, 51 minutes",
				Uptime_percent:  "100",
				Uptime_duration: 521490225893021,
				// Last_down:       "2025-01-10T22:27:46.640373Z",
				Last_down: lastDownTime,
			},
		}

		json.NewEncoder(w).Encode(statuses)
	}
}
