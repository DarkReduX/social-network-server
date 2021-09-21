package handler

import (
	"database/sql"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login godoc
// @Summary Log in user in account
// @Tags auth
// @Description Returns user token
// @Accept  json
// @Produce  json
// @Param profile body models.Profile true "Profile password and username"
// @Success 200 {string} token
// @Failure 400
// @Router /auth [post]
func (h AuthHandler) Login(c echo.Context) error {
	profile := models.Profile{}
	if err := c.Bind(&profile); err != nil {
		log.WithFields(log.Fields{
			"handler": "auth",
			"func":    "Login()",
		}).Errorf("Unable to bind data: %v", err)

		return err
	}
	token, err := h.authService.Login(c.Request().Context(), profile.Username, profile.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "auth",
			"func":    "Login()",
		}).Errorf("Unable to login user: %v", err)

		if err == sql.ErrNoRows {
			return echo.ErrBadRequest
		}
		return err
	}

	return c.JSON(http.StatusOK, token)
}

// Logout godoc
// @Summary Logout user
// @Tags auth
// @Description Delete user token
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Success 200
// @Failure 400
// @Router /auth [delete]
func (h AuthHandler) Logout(c echo.Context) error {
	username := c.Get("username").(string)
	if err := h.authService.Logout(c.Request().Context(), username); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
