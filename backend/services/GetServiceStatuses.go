package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "uptime/handlers"
	"uptime/models"
	"uptime/repository"
)

func GetServiceStatuses(db *sql.DB) ([]models.ServiceStatus, error) {
	servicesMap := map[string]string{
		"jaypatel": "https://jaypatel.link",
		"magicdot": "https://magicdot.jaypatel.link",
		"dev":      "https://dev.jaypatel.link",
		"app":      "https://app.jaypatel.link",
		"res":      "https://res.jaypatel.link",
		// "uptime":   "https://uptime.jaypatel.link",
	}

	var statuses []models.ServiceStatus
	for name, url := range servicesMap {
		status, uptime, uptimePercent, uptimeDuration, lastDown, err := CheckService(db, name, url)
		if err != nil {
			return nil, err
		}

		if status == "Down" {
			//fetch subscribers email id
			subscribers, err := repository.FetchAllSubscribers(db)
			if err != nil {
				log.Printf("Failed to fetch subscribers: %v", err)
				return nil, err
			}
			//send mails to subscribers
			for _, email := range subscribers {
				err := SendNotification(email, name)
				if err != nil {
					log.Printf("Failed to send email to %s:%v", email, err)
				}
			}
		}
		statuses = append(statuses, models.ServiceStatus{
			Name: name, Url: url, Status: status, Uptime: uptime,
			Uptime_percent: uptimePercent, Uptime_duration: uptimeDuration, Last_down: lastDown,
		})
	}

	return statuses, nil
}

func SendNotification(email string, servicename string) error {
	mailRequest := models.Message{
		To:           email,
		TemplateName: "notification",
		DataMap: map[string]any{
			"service": servicename,
		},
	}

	mailRequestBytes, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/send-mail", bytes.NewReader(mailRequestBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email, status code: %d", resp.StatusCode)
	}

	log.Printf("email sent successfully to %s", email)
	return nil
}
