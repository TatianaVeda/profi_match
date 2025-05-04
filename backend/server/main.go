package main

import (
	"log"
	"net/http"

	"backend/internal/database"
	"backend/internal/handlers"
)

func main() {
	err := database.InitDB("user=myuser password=mypassword dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
