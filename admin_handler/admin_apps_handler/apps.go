package admin_apps_handler

import (
	"net/http"
	"strconv"

	"github.com/MXslade/log_service_go/db/repo/app_repo"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	appRepo := app_repo.New()

	withEnvs := c.QueryParam("withEnvs")
	if len(withEnvs) > 0 {
		withEnvsParsed, err := strconv.ParseBool(withEnvs)
		if err != nil {
			c.JSON(http.StatusBadRequest, echo.Map{"error": err})
			return err
		}
		if withEnvsParsed {
			apps, err := appRepo.GetAllWithEnvs(c.Request().Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
				return err
			}

			return c.JSON(http.StatusOK, apps)
		}
	}

	apps, err := appRepo.GetAll(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		return err
	}
	return c.JSON(http.StatusOK, apps)
}

func Create(c echo.Context) error {
	return c.String(http.StatusCreated, "apps: Create")
}

func Show(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Show")
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Update")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Delete")
}
