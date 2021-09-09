package handler

import (
	"database/sql"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ProfileHandler struct {
	service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h ProfileHandler) Get(c echo.Context) error {
	id := c.QueryParam("id")

	profile, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "get",
		}).Errorf("Couldn't scan result: %v", err)

		if err == sql.ErrNoRows {
			return echo.ErrNotFound
		}

		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, profile)
}

func (h ProfileHandler) Create(c echo.Context) error {
	profile := models.Profile{}
	if err := c.Bind(&profile); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.service.Create(c.Request().Context(), profile); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "Create",
		}).Errorf("Couldn't create user profile: %v", err)

		return echo.ErrBadRequest
	}
	return c.NoContent(http.StatusOK)
}

func (h ProfileHandler) Update(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.ErrBadRequest
	}

	profile := models.Profile{}
	if err := c.Bind(&profile); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "Update",
		}).Errorf("Unable bind request body: %v", err)
	}

	if err := h.service.Update(c.Request().Context(), profile); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "Update",
		}).Errorf("Couldn't update profile: %v", err)

		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h ProfileHandler) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.ErrBadRequest
	}

	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "Update",
		}).Errorf("Couldn't mark profile as deleted: %v", err)

		return echo.ErrBadRequest
	}
	return c.NoContent(http.StatusOK)
}
