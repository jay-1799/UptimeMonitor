package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"uptime/models"
	"uptime/utils"
)

func UpdateLastDown(db *sql.DB, serviceName string) error {
	query := `
        UPDATE uptime_logs
        SET last_down = $1
        WHERE service_name = $2;`
	_, err := db.Exec(query, time.Now(), serviceName)
	if err != nil {
		return fmt.Errorf("failed to update last_down for service %s: %v", serviceName, err)
	}
	return nil
}

func FetchLastDown(db *sql.DB, serviceName string) (time.Time, error) {
	var lastDown time.Time
	query := `
		SELECT last_down 
		FROM uptime_logs
		WHERE service_name = $1;`

	err := db.QueryRow(query, serviceName).Scan(&lastDown)
	if err != nil {
		// panic(err)
		if err == sql.ErrNoRows {
			// No record found for the service
			return time.Time{}, fmt.Errorf("no downtime record found for service: %s", serviceName)
		}
		return time.Time{}, fmt.Errorf("failed to fetch last_down for service %s: %v", serviceName, err)
	}

	return lastDown, nil
}

/////////////////////////

func FetchUptimeData(db *sql.DB) []models.ServiceStatus {
	rows, err := db.Query("SELECT service_name, url, status,last_down  FROM uptime_logs ORDER BY timestamp DESC LIMIT 10")
	if err != nil {
		log.Printf("Failed to fetch uptime data: %v", err)
		return nil
	}
	defer rows.Close()

	var statuses []models.ServiceStatus
	for rows.Next() {
		var status models.ServiceStatus
		// var uptime time.Duration
		var last_down time.Time
		err := rows.Scan(&status.Name, &status.Url, &status.Status, &last_down)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		status.Uptime, _, _ = utils.Calculate_uptime(last_down)
		// status.last_down = last_down
		statuses = append(statuses, status)
	}
	return statuses
}

// Store service status in the database
// func storeServiceStatus(db *sql.DB, name, url, status string, uptime time.Duration) {
// 	_, err := db.Exec(
// 		"INSERT INTO uptime_logs (service_name, url, status, last_down) VALUES ($1, $2, $3, $4)",
// 		name, url, status, uptime,
// 	)
// 	if err != nil {
// 		log.Printf("Failed to store service status: %v", err)
// 	}
// }
