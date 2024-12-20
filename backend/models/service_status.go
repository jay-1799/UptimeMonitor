package models

import (
	"html/template"
	"time"
)

type ServiceStatus struct {
	Name            string        `json:"name"`
	Url             string        `json:"url"`
	Status          string        `json:"status"`
	Uptime          string        `json:"uptime"`
	Uptime_percent  string        `json:"uptime_percent"`
	Uptime_duration time.Duration `json:"uptime_duration"`
	Last_down       time.Time     `json:"last_down"`
}

type MailData struct {
	To      string
	From    string
	Subject string
	Content template.HTML
}
