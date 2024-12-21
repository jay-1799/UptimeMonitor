package handlers

import (
	"encoding/json"
	"net/http"
	"uptime/mail"
)

type MailHandler struct {
	Mailer *mail.Mail
}

func (mh *MailHandler) SendMail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	plainText := req.Message
	htmlText := "<p>" + req.Message + "</p>"

	err = mh.Mailer.Send(req.To, req.Subject, plainText, htmlText)
	if err != nil {
		// panic(err)
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
