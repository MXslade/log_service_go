package route

import (
	"github.com/MXslade/log_service_go/handler/apps_handler"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	v1Group := e.Group("/api/v1")
	{
		appsGroup := v1Group.Group("/apps")
		{
			appsGroup.GET("", apps_handler.Index)
			appsGroup.POST("", apps_handler.Create)
			appsGroup.GET("/:id", apps_handler.Show)
			appsGroup.PUT("/:id", apps_handler.Update)
			appsGroup.DELETE("/:id", apps_handler.Delete)
		}
	}
}
