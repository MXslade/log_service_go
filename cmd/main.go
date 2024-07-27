package main

import (
	"log"

	"github.com/MXslade/log_service_go/app"
	"github.com/MXslade/log_service_go/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Initializing db connection pool")
	db.InitDBPool()
	defer db.CloseDBPool()
	db.RunMigrations()

	log.Println("Initializing echo")
	e := echo.New()
    app.InitApp(e)
    app.StartApp(e)
}
