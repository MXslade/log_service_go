package app

import (
	"fmt"
	"log"
	"os"

	"github.com/MXslade/log_service_go/config"
	"github.com/MXslade/log_service_go/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitApp(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogMethod:   true,
		LogError:    true,
		LogProtocol: true,
		LogLatency:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("%v %v %v %v %v\n\n", v.Protocol, v.Method, v.URI, v.Status, v.Latency)
			if v.Error != nil {
				log.Printf("Error: %v\n\n", v.Error)
			}
			return nil
		},
	}))

	e.Use(middleware.Recover())

	route.SetUpRoutes(e)
}

func StartApp(e *echo.Echo) {
	appPort, ok := os.LookupEnv("APP_PORT")
	if ok {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", appPort)))
	} else {
		log.Printf("no APP_PORT env value is set, using default: %v\n", config.DefaultAppPort)
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.DefaultAppPort)))
	}
}
