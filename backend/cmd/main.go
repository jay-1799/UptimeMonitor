package main

import (
	"log"
	"net/http"
	"uptime/db"
	"uptime/handlers"
	"uptime/middlewares"

	_ "github.com/lib/pq"
)

func main() {
	db := db.InitDB()
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler(db))
	mux.HandleFunc("/history", handlers.HistoryHandler(db))
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", middlewares.CorsMiddleware(mux)))
}
