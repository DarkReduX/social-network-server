package handler

import (
	"database/sql"
	"fmt"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

type ProfileHandler struct {
	service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// Get godoc
// @Summary Get user profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param id query string true "User profile ID"
// @Success 200 {object} models.Profile
// @Failure 400
// @Failure 500
// @Router /profile [get]
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

	profile.Password = ""
	return c.JSON(http.StatusOK, profile)
}

// Create godoc
// @Summary Create user profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param profile body models.Profile true "User profile fields"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /profile [post]
func (h ProfileHandler) Create(c echo.Context) error {
	profile := models.Profile{}
	if err := c.Bind(&profile); err != nil {
		return echo.ErrBadRequest
	}
	profile.UUID = uuid.New().String()

	if err := h.service.Create(c.Request().Context(), profile); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "Create",
		}).Errorf("Couldn't create user profile: %v", err)

		return echo.ErrBadRequest
	}
	return c.NoContent(http.StatusOK)
}

// Update godoc
// @Summary Update user profile info
// @Tags profile
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param id query int true "User id"
// @Param profile body models.Profile true "User profile fields"
// @Success 200
// @Failure 400
// @Router /profile [put]
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

// Delete godoc
// @Summary Set user profile as deleted
// @Tags profile
// @Accept  json
// @Produce  json
// @Security BearerToken
// @Param id query int true "User profile ID"
// @Success 200
// @Failure 400
// @Router /profile [delete]
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

// UploadUserAvatar godoc
// @Summary Set user profile avatar
// @Tags profile
// @Accept  mpfd
// @Produce  json
// @Security BearerToken
// @Param img-path path string true "Server img path"
// @Param avatar formData file true "Upload image"
// @Success 200
// @Failure 400
// @Router /profile/{img-path} [put]
func (h ProfileHandler) UploadUserAvatar(c echo.Context) error {
	username := c.Get("username").(string)
	filePath := c.Param("img-path")
	filePath = strings.ReplaceAll(filePath, "%2F", "/")

	currentPrjPath, err := os.Getwd()
	if err != nil {
		log.Errorf("Couldn't get current project path: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if !strings.Contains(filePath, currentPrjPath) {
		return c.JSON(http.StatusBadRequest, "Incorrect path")
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "UploadUserAvatar",
		}).Errorf("Couldn't get file form: %v", err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "UploadUserAvatar",
		}).Errorf("Couldn't open file: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	fullPath := fmt.Sprintf("%s/%s", filePath, file.Filename)
	dst, err := os.Create(fullPath)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "UploadUserAvatar",
		}).Errorf("Couldn't create file: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}
	defer func() {
		if err := dst.Close(); err != nil {
			log.WithFields(log.Fields{
				"handler": "profile",
				"func":    "UploadUserAvatar",
			}).Errorf("Couldn't close created file: %v", err)
		}
	}()

	if _, err := io.Copy(dst, src); err != nil {
		log.WithFields(log.Fields{
			"handler": "profile",
			"func":    "UploadUserAvatar",
		}).Errorf("Couldn't copy sent file: %v", err)
	}

	userProfile, err := h.service.Get(c.Request().Context(), username)
	userProfile.AvatarLink = &fullPath
	if err != nil {
		return err
	}

	if err := h.service.Update(c.Request().Context(), *userProfile); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Avatar %s uploaded successfully for user %s", file.Filename, username))
}
