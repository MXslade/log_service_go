package main

import (
	"log"

	"github.com/MXslade/log_service_go/route"
	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags)

	log.Println("Initializing echo")
	e := echo.New()

	route.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":1313"))
}
