package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"uptime/repository"
)

func SubscriberHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var requestData struct {
			EmailID string `json:"emailID"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		err := repository.AddSubscriber(db, requestData.EmailID)
		if err != nil {
			http.Error(w, "Failed to add subscriber", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"subscriber added successfully}`))
	}
}
