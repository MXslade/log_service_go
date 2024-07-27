package main

import (
	"flag"
	"log"

	"github.com/MXslade/log_service_go/admin_cli"
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

	mode := flag.String("mode", "server", "Defines the mode of the application. \n\"server\": when you want to run a server.\n \"admin_cli\": when you want to run admin cli(you can create admin user there)\n.")
	flag.Parse()
	if mode == nil {
		log.Fatal("Mode is undefined.")
	}

	log.Printf("Mode: %v\n", *mode)

	if *mode == "admin_cli" {
		log.Println("Running ADMIN CLI app")
		admin_cli.Start()
	} else if *mode == "server" {
		log.Println("Running SERVER app")

		log.Println("Initializing echo")
		e := echo.New()
		app.InitApp(e)
		app.StartApp(e)
	}
}
