package mail

import (
	"bytes"
	"log"
	"net/smtp"
	"strconv"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	FromName    string
	FromAddress string
}

func (m *Mail) Send(to, subject, plainText, htmlText string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	msg := bytes.Buffer{}
	msg.WriteString("Subject: " + subject + "\n")
	msg.WriteString("To: " + to + "\n")
	msg.WriteString("MIME-Version: 1.0\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\n\n")
	msg.WriteString("htmlText")

	// iPort := strconv.Itoa(m.Port)
	// host := m.Host + ":" + iPort
	err := smtp.SendMail(
		m.Host+":"+strconv.Itoa(m.Port),
		auth,
		m.FromAddress,
		[]string{to},
		msg.Bytes(),
	)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}
	log.Println("Email sent successfully to", to)
	return nil
}
