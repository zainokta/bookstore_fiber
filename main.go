package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zainokta/bookstore_fiber/database"
	"github.com/zainokta/bookstore_fiber/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.InitDatabase().Client
	app := routers.Router(db)
	defer db.Close()
	
	app.Listen(8000)
}
