package admin_auth_handler

import (
	"errors"
	"net/http"

	"github.com/MXslade/log_service_go/service/jwt_service"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var data loginData

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"error": err})
		return err
	}

	token, err := jwt_service.Login(
		c.Request().Context(),
		jwt_service.LoginData{
			Username: data.Username,
			Password: data.Password,
		},
	)
	if err != nil {
		if errors.Is(err, echo.ErrUnauthorized) {
			return err
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
