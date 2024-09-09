package admin_apps_handler

import (
	"context"
	"net/http"
	"strconv"

	model_app "github.com/MXslade/log_service_go/model/app_model"
	"github.com/labstack/echo/v4"
)

type appRepo interface {
	GetAll(ctx context.Context) ([]*model_app.AppModel, error)
	GetAllWithEnvs(ctx context.Context) ([]*model_app.AppWithEnvs, error)
}

type appsHandler struct {
	appRepo appRepo
}

func New(appRepo appRepo) *appsHandler {
	return &appsHandler{appRepo: appRepo}
}

func (h *appsHandler) Index(c echo.Context) error {
	withEnvs := c.QueryParam("withEnvs")
	if len(withEnvs) > 0 {
		withEnvsParsed, err := strconv.ParseBool(withEnvs)
		if err != nil {
			c.JSON(http.StatusBadRequest, echo.Map{"error": err})
			return err
		}
		if withEnvsParsed {
			apps, err := h.appRepo.GetAllWithEnvs(c.Request().Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
				return err
			}

			return c.JSON(http.StatusOK, apps)
		}
	}

	apps, err := h.appRepo.GetAll(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		return err
	}
	return c.JSON(http.StatusOK, apps)
}

func (h *appsHandler) Create(c echo.Context) error {
	return c.String(http.StatusCreated, "apps: Create")
}

func (h *appsHandler) Show(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Show")
}

func (h *appsHandler) Update(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Update")
}

func (h *appsHandler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "apps: Delete")
}
