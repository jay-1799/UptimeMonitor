package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"uptime/repository"
)

func StatusPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var formData struct {
			EmailId     string `json:"emailID"`
			Username    string `json:"username"`
			Projectlink string `json:"projectLink"`
		}
		if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}
		_, _ = repository.CreateStatusPage(db, formData.EmailId, formData.Username, formData.Projectlink)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"status page created"}`))

	}
}
