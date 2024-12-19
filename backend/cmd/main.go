package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type ServiceStatus struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

type MailData struct {
	To      string
	From    string
	Subject string
	Content template.HTML
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

// Initialize database connection
func initDB() *sql.DB {
	connStr := "postgres://postgres:password@postgres:5432/uptime_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

// Store service status in the database
func storeServiceStatus(db *sql.DB, name, url, status string, uptime time.Duration) {
	_, err := db.Exec(
		"INSERT INTO uptime_logs (service_name, url, status, last_down) VALUES ($1, $2, $3, $4)",
		name, url, status, uptime,
	)
	if err != nil {
		log.Printf("Failed to store service status: %v", err)
	}
}

// Retrieve uptime data from the database
func fetchUptimeData(db *sql.DB) []ServiceStatus {
	rows, err := db.Query("SELECT service_name, url, status,  FROM uptime_logs ORDER BY timestamp DESC LIMIT 10")
	if err != nil {
		log.Printf("Failed to fetch uptime data: %v", err)
		return nil
	}
	defer rows.Close()

	var statuses []ServiceStatus
	for rows.Next() {
		var status ServiceStatus
		// var uptime time.Duration
		var last_down time.Time
		err := rows.Scan(&status.Name, &status.Url, &status.Status, &last_down)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		status.Uptime = calculate_uptime(last_down)
		// status.last_down = last_down
		statuses = append(statuses, status)
	}
	return statuses
}

func calculate_uptime(last_down time.Time) string {
	current_time := time.Now()

	uptime_duration := current_time.Sub(last_down)

	// Format the duration to a string (e.g., "2 days 3 hours")
	days := int(uptime_duration.Hours()) / 24
	hours := int(uptime_duration.Hours()) % 24
	minutes := int(uptime_duration.Minutes()) % 60

	return formatUptime(days, hours, minutes)
}

func formatUptime(days, hours, minutes int) string {
	if days > 0 {
		return timeString(days, "day") + ", " + timeString(hours, "hour") + ", " + timeString(minutes, "minute")
	}
	if hours > 0 {
		return timeString(hours, "hour") + ", " + timeString(minutes, "minute")
	}
	return timeString(minutes, "minute")
}

func timeString(value int, unit string) string {
	if value == 1 {
		return "1 " + unit
	}
	// return time.Duration(value).String()
	return strconv.Itoa(value) + " " + unit + "s"
}

func checkService(db *sql.DB, serviceName, url string) (string, string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		// todo: function to update downtime timestamp
		if err := updateLastDown(db, serviceName); err != nil {
			log.Printf("Error updating downtime for %s: %v", serviceName, err)
		}
		return "Down", "0", nil
	}
	defer resp.Body.Close()
	last_down, err := fetchLastDown(db, serviceName)
	if err != nil {
		return "", "", err
	}
	uptime := calculate_uptime(last_down)
	return "Up", uptime, nil
	// return "Up", "10 days", nil
}

func fetchLastDown(db *sql.DB, serviceName string) (time.Time, error) {
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

func updateLastDown(db *sql.DB, serviceName string) error {
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

// func statusHandler(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-type", "application/json")

// 	services := map[string]string{
// 		"jaypatel": "https://jaypatel.link",
// 		"magicdot": "https://magicdot.jaypatel.link",
// 		"dev":      "https://dev.jaypatel.link",
// 		"app":      "https://app.jaypatel.link",
// 		"res":      "https://res.jaypatel.link",
// 		"uptime":   "https://uptime.jaypatel.link",
// 	}

//		var statuses []ServiceStatus
//		for name, url := range services {
//			status := checkService(url)
//			uptime := "99.1"
//			statuses = append(statuses, ServiceStatus{Name: name, Url: url, Status: status, Uptime: uptime})
//		}
//		json.NewEncoder(w).Encode(statuses)
//	}
func statusHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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
			status, uptime, _ := checkService(db, name, url)

			statuses = append(statuses, ServiceStatus{Name: name, Url: url, Status: status, Uptime: uptime})

			// Store in the database
			// storeServiceStatus(db, name, url, status, uptime)
		}

		json.NewEncoder(w).Encode(statuses)
	}
}
func historyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		statuses := fetchUptimeData(db)
		json.NewEncoder(w).Encode(statuses)
	}
}

func main() {
	db := initDB()
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/status", statusHandler(db))
	mux.HandleFunc("/history", historyHandler(db))
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
