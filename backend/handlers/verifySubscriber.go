package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"uptime/repository"
)

func VerifySubscriberHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "token is required", http.StatusBadRequest)
			return
		}
		email, err := repository.VerifySubscriber(db, token)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{message:"email %s successfully verified"}`, email)))
	}
}
