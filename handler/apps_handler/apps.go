package apps_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Index")
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
