package services

import (
	"database/sql"
	"uptime/models"
)

func GetServiceStatuses(db *sql.DB) ([]models.ServiceStatus, error) {
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
		status, uptime, uptimePercent, uptimeDuration, lastDown, err := CheckService(db, name, url)
		if err != nil {
			return nil, err
		}

		if status == "Down" {
			//fetch subscribers email id

			// subscribers := repository.FetchSubscribers()

			//send mails to subscribers

		}
		statuses = append(statuses, models.ServiceStatus{
			Name: name, Url: url, Status: status, Uptime: uptime,
			Uptime_percent: uptimePercent, Uptime_duration: uptimeDuration, Last_down: lastDown,
		})
	}

	return statuses, nil
}
