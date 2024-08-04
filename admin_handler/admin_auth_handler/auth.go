package admin_auth_handler

import (
	"errors"
	"net/http"

	"github.com/MXslade/log_service_go/db/repo/admin_repo"
	"github.com/MXslade/log_service_go/service/auth_service"
	"github.com/MXslade/log_service_go/service/jwt_service"
	"github.com/jackc/pgx/v5"
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

	adminRepo, err := admin_repo.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		return err
	}

	admin, err := adminRepo.GetByUsername(c.Request().Context(), data.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, echo.Map{"error": err})
			return err
		} else {
			c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
			return err
		}
	}

	authService, err := auth_service.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		return err
	}

	ok := authService.VerifyHash(data.Password, admin.Password)
	if !ok {
		return echo.ErrUnauthorized
	}

	token, err := jwt_service.CreateToken(
		c.Request().Context(),
		admin,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
