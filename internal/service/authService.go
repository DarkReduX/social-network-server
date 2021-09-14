package service

import (
	"context"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
