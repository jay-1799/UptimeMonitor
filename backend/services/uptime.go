package services

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"uptime/repository"
	"uptime/utils"
)

func CheckService(db *sql.DB, serviceName, url string) (string, string, string, time.Duration, time.Time, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		// todo: function to update downtime timestamp
		if err := repository.UpdateLastDown(db, serviceName); err != nil {
			log.Printf("Error updating downtime for %s: %v", serviceName, err)
		}
		return "Down", "0", "", 0 * time.Second, time.Now(), nil
	}
	defer resp.Body.Close()
	last_down, err := repository.FetchLastDown(db, serviceName)
	if err != nil {
		return "", "", "", 0 * time.Second, time.Now(), err
	}
	uptime, uptime_percent, uptime_duration := utils.Calculate_uptime(last_down)
	return "Up", uptime, uptime_percent, uptime_duration, last_down, nil
	// return "Up", "10 days", nil
}
