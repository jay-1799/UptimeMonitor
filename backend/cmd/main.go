package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"uptime/db"
	"uptime/handlers"
	"uptime/mail"
	"uptime/middlewares"
	"uptime/services"

	_ "github.com/lib/pq"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	// port, err := strconv.Atoi(os.Getenv("PORT"))
	// if err != nil {
	// 	log.Fatalf("failed to parse int: %v", err)
	// }

	mailConfig := &mail.Mail{
		Domain: os.Getenv("MAIL_DOMAIN"),
		Host:   os.Getenv("MAIL_HOST"),
		// Host: "mailhog",
		Port: atoi(os.Getenv("MAIL_PORT")),
		// 1025,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
	}

	scheduler := services.NewScheduler()
	scheduler.AddJob(90*time.Minute, func() {
		log.Println("Performing task")
		services.PerformPeriodicTask(db)
	})
	scheduler.StartScheduler()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler(db))
	mux.HandleFunc("/history", handlers.HistoryHandler(db))
	mux.HandleFunc("/add-subscriber", handlers.SubscriberHandler(db))
	mux.HandleFunc("/verify-subscriber", handlers.VerifySubscriberHandler(db))

	mailHandler := &handlers.MailHandler{Mailer: mailConfig}
	mux.HandleFunc("/send-mail", mailHandler.SendMail)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", middlewares.CorsMiddleware(mux)))
}

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v", err)
	}
	return val
}
