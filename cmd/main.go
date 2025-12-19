package main

import (
	"cli-app/database"
	"cli-app/ui"
	"log"
	"os"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if cerr := db.Close(); cerr != nil {
			log.Printf("Error closing database connection: %v", cerr)
		}
	}()

	if err := ui.NewApp(db).Run(); err != nil {
		log.Printf("Application error: %v", err)
		os.Exit(1)
	}
}
