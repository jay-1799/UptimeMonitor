package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"
	"uptime/mail"
	"uptime/models"

	"github.com/vanng822/go-premailer/premailer"
)

type MailHandler struct {
	Mailer *mail.Mail
}

func (mh *MailHandler) SendMail(w http.ResponseWriter, r *http.Request) {
	// var req struct {
	// 	To      string `json:"to"`
	// 	Subject string `json:"subject"`
	// 	Message string `json:"message"`
	// 	Service string `json:"service`
	// }
	var req models.Message

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// req.DataMap = map[string]any{
	// 	"service": req.Service,
	// }
	plainTextTemplate := "/app/mail/templates/" + req.TemplateName + ".plain.gohtml"
	htmlTemplate := "/app/mail/templates/" + req.TemplateName + ".html.gohtml"

	plainText, err := mh.buildPlainTextMessage(plainTextTemplate, req.DataMap)
	if err != nil {
		panic(err)
		// http.Error(w, "Failed to build plain text message", http.StatusInternalServerError)
		// return
	}
	htmlText, err := mh.buildHTMLMessage(htmlTemplate, req.DataMap)
	if err != nil {
		http.Error(w, "Failed to build HTML message", http.StatusInternalServerError)
		return
	}
	// htmlText := "<p>" + req.Message + "</p>"

	err = mh.Mailer.Send(req.To, req.Subject, plainText, htmlText)
	if err != nil {
		// panic(err)
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

func (mh *MailHandler) buildHTMLMessage(templateToRender string, data map[string]any) (string, error) {
	// templateToRender := "/app/mail/templates/mail.html.gohtml"
	// "../mail/templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}
	// msg.DataMap = map[string]any{
	// 	"service": msg.Service,
	// }

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = mh.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (mh *MailHandler) buildPlainTextMessage(templateToRender string, data map[string]any) (string, error) {
	// templateToRender := "/app/mail/templates/mail.plain.gohtml"
	// "../mail/templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (mh *MailHandler) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}
	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}
	html, err := prem.Transform()
	if err != nil {
		return "", err
	}
	return html, nil
}
