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

type Message struct {
	From         string         `json:"from"`
	FromName     string         `json:"fromname"`
	To           string         `json:"to"`
	Subject      string         `json:"subject"`
	Service      string         `json:"service"`
	TemplateName string         `json:"template_name"`
	DataMap      map[string]any `json:"data_map"`
}
