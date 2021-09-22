package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DarkReduX/social-network-server/internal/models"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r ProfileRepository) GetByUUID(ctx context.Context, id string) (*models.Profile, error) {
	profile := models.Profile{}
	query := `select * from get_profile($1)`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&profile.UUID, &profile.Username, &profile.Password, &profile.AvatarLink, &profile.LastActivity, &profile.CreatedAt, &profile.CreatedFromIp, &profile.DeletedAt, &profile.IsActivate)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r ProfileRepository) GetByName(ctx context.Context, username string) (*models.Profile, error) {
	profile := models.Profile{}
	query := `select * from get_profile_by_name($1)`
	err := r.db.QueryRowContext(ctx, query, username).Scan(&profile.UUID, &profile.Username, &profile.Password, &profile.AvatarLink, &profile.LastActivity, &profile.CreatedAt, &profile.CreatedFromIp, &profile.DeletedAt, &profile.IsActivate)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r ProfileRepository) Create(ctx context.Context, profile models.Profile) error {
	query := `select * from create_profile($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.db.ExecContext(ctx, query, profile.UUID, profile.Username, profile.Password, profile.AvatarLink, profile.LastActivity, profile.CreatedAt, profile.CreatedFromIp)
	return err
}

func (r ProfileRepository) Update(ctx context.Context, profile models.Profile) error {
	query := `select * from update_profile($1,$2,$3,$4,$5,$6,$7)`
	res, err := r.db.ExecContext(ctx, query, profile.UUID, profile.Username, profile.Password, profile.AvatarLink, profile.LastActivity, profile.DeletedAt, profile.IsActivate)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if rows == 0 {
		return errors.New("Zero rows affected ")
	}
	return err
}

func (r ProfileRepository) Delete(ctx context.Context, id string) error {
	query := `select * from delete_profile($1)`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return errors.New("Zero rows affected ")
	}
	return err
}
