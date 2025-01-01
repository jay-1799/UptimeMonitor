package handlers

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"uptime/models"
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
			// panic(err)
		}
		//todo generating token
		token, err := generateToken()
		if err != nil {
			http.Error(w, "failed to generate verification token", http.StatusInternalServerError)
			return
			// panic(err)
		}
		// save the token to database
		msg, err := repository.AddSubscriberToken(db, requestData.EmailID, token)
		if err != nil {
			http.Error(w, "Failed to add subscriber token", http.StatusInternalServerError)
			return
		}
		if msg == fmt.Sprintf("email %s is already verified", requestData.EmailID) {
			log.Printf("Email %s is already verified. Skipping email sending.", requestData.EmailID)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"user is already verified, no email sent"}`))
			return
		}

		//send the verification link
		email := requestData.EmailID
		log.Print(requestData.EmailID)
		// verificationLink := fmt.Sprintf("http://localhost:8080/verify-subscriber?token=%s", token)
		verificationLink := fmt.Sprintf("https://status.jaypatel.link/verify-subscriber?token=%s", token)
		mailRequest := models.Message{
			To:           email,
			TemplateName: "activation",
			DataMap: map[string]any{
				"activation_link": verificationLink,
			},
		}
		mailRequestBytes, err := json.Marshal(mailRequest)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "http://localhost:8080/send-mail", bytes.NewReader(mailRequestBytes))
		if err != nil {
			http.Error(w, "failed to create request to SendMail", http.StatusInternalServerError)
			return
			// panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "failed to send verification email", http.StatusInternalServerError)
			return
			// panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "failed to send verification email", resp.StatusCode)
			return
			// panic(err)
		}
		log.Printf("email sent successfully to %s", email)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"verification email sent"}`))
	}
}

func generateToken() (string, error) {
	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
