package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MXslade/log_service_go/config"
	"github.com/MXslade/log_service_go/db"
	"github.com/MXslade/log_service_go/route"
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

	log.Println("Initializing echo")
	e := echo.New()

	route.SetUpRoutes(e)

	appPort, ok := os.LookupEnv("APP_PORT")
	if ok {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", appPort)))
	} else {
		log.Printf("no APP_PORT env value is set, using default: %v\n", config.DefaultAppPort)
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.DefaultAppPort)))
	}
}
