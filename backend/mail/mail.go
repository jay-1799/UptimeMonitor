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

	// iPort := strconv.Itoa(m.Port)
	// host := m.Host + ":" + iPort

	// err := smtp.SendMail(
	// 	m.Host+":"+strconv.Itoa(m.Port),
	// 	auth,
	// 	m.FromAddress,
	// 	[]string{to},
	// 	msg.Bytes(),
	// )
	//
	log.Printf("Attempting to connect to SMTP server %s on port %d", m.Host, m.Port)
	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
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
	// if len(msg.Attachments) > 0 {
	// 	for _, x := range msg.Attachments {
	// 		email.AddAttachment(x)

	// 	}
	// }
	err = email.Send(smtpClient)
	if err != nil {
		log.Println("Error2 sending email:", err)
		return err
	}
	return nil
	//

	// if err != nil {
	// 	log.Println("Error sending email:", err)
	// 	return err
	// }
	// log.Println("Email sent successfully to", to)
	// return nil
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
