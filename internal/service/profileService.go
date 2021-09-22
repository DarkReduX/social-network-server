package service

import (
	"context"
	"fmt"
	"github.com/DarkReduX/social-network-server/internal/models"
	"github.com/DarkReduX/social-network-server/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ProfileService struct {
	repository *repository.ProfileRepository
}

func NewProfileService(repository *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repository: repository}
}

func (s ProfileService) Get(ctx context.Context, username string) (*models.Profile, error) {
	profile, err := s.repository.GetByName(ctx, username)
	return profile, err
}

func (s ProfileService) Create(ctx context.Context, profile models.Profile) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(profile.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	profile.Password = string(passwordHash)
	profile.CreatedAt = &time.Time{}
	*profile.CreatedAt = time.Now()
	profile.LastActivity = profile.CreatedAt

	return s.repository.Create(ctx, profile)
}

func (s ProfileService) Update(ctx context.Context, profile models.Profile) error {
	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s", profile.Password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profile.Password = string(passwordHash)

	return s.repository.Update(ctx, profile)
}

func (s ProfileService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
