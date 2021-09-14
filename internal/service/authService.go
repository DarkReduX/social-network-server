package service

import (
	"context"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type AuthService struct {
	profileRepository *repository.ProfileRepository
	authRepository    *repository.AuthRepository
}

func NewAuthService(profileRepository *repository.ProfileRepository, authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		profileRepository: profileRepository,
		authRepository:    authRepository,
	}
}

func (s AuthService) Login(ctx context.Context, username string, password string) (*string, error) {
	profile, err := s.profileRepository.Get(ctx, username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		return nil, echo.ErrUnauthorized
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
	}).SignedString([]byte("user"))
	err = s.authRepository.WriteToken(ctx, username, token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s AuthService) Logout(ctx context.Context, username string) error {
	return s.authRepository.WriteToken(ctx, username, "")
}

func (s *AuthService) AuthenticateTokenMiddleware(config middleware.JWTConfig) echo.MiddlewareFunc {
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

			redisToken, err := s.authRepository.GetToken(c.Request().Context(), claim.Username)
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
