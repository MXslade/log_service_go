package admin_auth_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/MXslade/log_service_go/model"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type adminRepo interface {
	GetByUsername(ctx context.Context, username string) (*model.AdminModel, error)
}

type authService interface {
	VerifyHash(toHash string, actualHash string) bool
}

type jwtService interface {
	CreateToken(ctx context.Context, admin *model.AdminModel) (string, error)
}

type authHandler struct {
	adminRepo   adminRepo
	authService authService
	jwtService  jwtService
}

func New(adminRepo adminRepo, authService authService, jwtService jwtService) *authHandler {
	return &authHandler{
		adminRepo:   adminRepo,
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *authHandler) Login(c echo.Context) error {
	var data loginData

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"error": err})
		return err
	}

	admin, err := h.adminRepo.GetByUsername(c.Request().Context(), data.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, echo.Map{"error": err})
			return err
		} else {
			c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
			return err
		}
	}

	ok := h.authService.VerifyHash(data.Password, admin.Password)
	if !ok {
		return echo.ErrUnauthorized
	}

	token, err := h.jwtService.CreateToken(
		c.Request().Context(),
		admin,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
