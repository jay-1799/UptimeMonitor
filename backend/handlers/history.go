package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"uptime/repository"
)

func HistoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		statuses := repository.FetchUptimeData(db)
		json.NewEncoder(w).Encode(statuses)
	}
}
