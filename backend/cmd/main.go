package main

import (
	"log"
	"net/http"
	"os"
	"uptime/db"
	"uptime/handlers"
	"uptime/mail"
	"uptime/middlewares"

	_ "github.com/lib/pq"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	// port, err := strconv.Atoi(os.Getenv("PORT"))
	// if err != nil {
	// 	log.Fatalf("failed to parse int: %v", err)
	// }

	mailConfig := mail.Mail{
		Domain: os.Getenv("MAIL_DOMAIN"),
		Host:   os.Getenv("MAIL_HOST"),
		Port:   1025,
		// atoi(os.Getenv("MAIL_PORT")),
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler(db))
	mux.HandleFunc("/history", handlers.HistoryHandler(db))

	mailHandler := &handlers.MailHandler{Mailer: &mailConfig}
	mux.HandleFunc("/send-mail", mailHandler.SendMail)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", middlewares.CorsMiddleware(mux)))
}

// func atoi(s string) int {
// 	val, err := strconv.Atoi(s)
// 	if err != nil {
// 		log.Fatalf("failed to parse int: %v", err)
// 	}
// 	return val
// }
