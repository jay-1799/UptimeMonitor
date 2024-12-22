package utils

import (
	"strconv"
	"time"
)

func Calculate_uptime(lastDown time.Time) (string, string, time.Duration) {
	currentTime := time.Now()
	uptimeDuration := currentTime.Sub(lastDown)

	days := int(uptimeDuration.Hours()) / 24
	hours := int(uptimeDuration.Hours()) % 24
	minutes := int(uptimeDuration.Minutes()) % 60

	uptimePercent := CalculateUptimePercent(days)
	return FormatUptime(days, hours, minutes), uptimePercent, uptimeDuration
}

func CalculateUptimePercent(days int) string {
	if days < 30 {
		// return strconv.Itoa((days / 30) * 100)
		uptimePercent := (float64(days) / 30.0) * 100
		return strconv.Itoa(int(uptimePercent))
	}
	return "100"
}

func FormatUptime(days, hours, minutes int) string {
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
	return strconv.Itoa(value) + " " + unit + "s"
}
