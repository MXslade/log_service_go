package admin_auth_handler

import (
	"log"
	"net/http"

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

	log.Printf("%+v: ", data)

	return c.JSON(http.StatusOK, echo.Map{"token": "token"})
}
