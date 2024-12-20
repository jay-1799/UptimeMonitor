package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"uptime/models"
	"uptime/services"
)

func StatusHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		servicesMap := map[string]string{
			"jaypatel": "https://jaypatel.link",
			"magicdot": "https://magicdot.jaypatel.link",
			"dev":      "https://dev.jaypatel.link",
			"app":      "https://app.jaypatel.link",
			"res":      "https://res.jaypatel.link",
			"uptime":   "https://uptime.jaypatel.link",
		}

		var statuses []models.ServiceStatus
		for name, url := range servicesMap {
			status, uptime, uptimePercent, uptimeDuration, lastDown, _ := services.CheckService(db, name, url)
			statuses = append(statuses, models.ServiceStatus{
				Name: name, Url: url, Status: status, Uptime: uptime,
				Uptime_percent: uptimePercent, Uptime_duration: uptimeDuration, Last_down: lastDown,
			})
		}

		json.NewEncoder(w).Encode(statuses)
	}
}
