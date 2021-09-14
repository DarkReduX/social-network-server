package middleware

import (
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthenticateToken(config middleware.JWTConfig, repository *repository.AuthRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, err := jwt.ParseWithClaims(
				strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "),
				&models.CustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return config.SigningKey, nil
				},
			)

			if err != nil {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"unable to parse token",
				)
			}

			claim, ok := token.Claims.(*models.CustomClaims)
			if !ok {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"couldn't get claims from token",
				)
			}

			redisToken, err := repository.GetToken(c.Request().Context(), claim.Username)
			if err != nil {
				log.WithFields(log.Fields{
					"service": "auth",
					"func":    "AuthenticateToken()",
				}).Errorf("Some error on get token from redis ")
			}

			if redisToken != token.Raw {
				return echo.ErrUnauthorized
			}

			c.Set("username", claim.Username)
			return next(c)
		}
	}
}
