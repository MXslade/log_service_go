package route

import (
	"github.com/MXslade/log_service_go/admin_handler/admin_apps_handler"
	"github.com/MXslade/log_service_go/admin_handler/admin_auth_handler"
	"github.com/MXslade/log_service_go/handler/apps_handler"
	"github.com/MXslade/log_service_go/service/jwt_service"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")

	v1ApiGroup := v1Group.Group("/api")
	{
		appsGroup := v1ApiGroup.Group("/apps")
		{
			appsGroup.GET("", apps_handler.Index)
			appsGroup.POST("", apps_handler.Create)
			appsGroup.GET("/:id", apps_handler.Show)
			appsGroup.PUT("/:id", apps_handler.Update)
			appsGroup.DELETE("/:id", apps_handler.Delete)
		}
	}

	v1AdminAuthGroup := v1Group.Group("/admin/auth")
	{
		v1AdminAuthGroup.POST("/login", admin_auth_handler.Login)
	}

	v1AdminApiGroup := v1Group.Group("/admin/api")
	v1AdminApiGroup.Use(echojwt.WithConfig(jwt_service.CreateConfig()))
	{
		appsGroup := v1AdminApiGroup.Group("/apps")
		{
			appsGroup.GET("", admin_apps_handler.Index)
			appsGroup.POST("", admin_apps_handler.Create)
			appsGroup.GET("/:id", admin_apps_handler.Show)
			appsGroup.PUT("/:id", admin_apps_handler.Update)
			appsGroup.DELETE("/:id", admin_apps_handler.Delete)
		}
	}
}
