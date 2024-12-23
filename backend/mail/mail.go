package mail

import (
	"bytes"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromName    string
	FromAddress string
}

func (m *Mail) Send(to, subject, plainText, htmlText string) error {
	// auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	msg := bytes.Buffer{}
	msg.WriteString("Subject: " + subject + "\n")
	msg.WriteString("To: " + to + "\n")
	msg.WriteString("MIME-Version: 1.0\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\n\n")
	msg.WriteString("htmlText")

	log.Printf("Attempting to connect to SMTP server %s on port %d with %s encryption", m.Host, m.Port, m.Encryption)
	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	// server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println("Error1 sending email:", err)
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(m.FromAddress).
		AddTo(to).
		SetSubject(subject)

	email.SetBody(mail.TextPlain, plainText)
	email.AddAlternative(mail.TextHTML, htmlText)

	err = email.Send(smtpClient)
	if err != nil {
		log.Println("Error2 sending email:", err)
		return err
	}
	return nil

}

// //////////
func (m *Mail) getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
