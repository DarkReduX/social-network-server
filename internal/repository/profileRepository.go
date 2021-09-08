package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DarkReduX/social-network-server/internal/models"
	"time"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r ProfileRepository) Get(ctx context.Context, id string) (*models.Profile, error) {
	profile := models.Profile{}
	err := r.db.QueryRowContext(ctx, `select * from profile where username = $1`, id).Scan(&profile.Username, &profile.Password, &profile.AvatarLink, &profile.LastActivity, &profile.CreatedAt, &profile.CreatedFromIp, &profile.DeletedAt, &profile.IsActivate)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r ProfileRepository) Create(ctx context.Context, profile models.Profile) error {
	_, err := r.db.ExecContext(ctx, `insert into profile values ($1,$2,$3,$4,$5,$6,$7,$8)`, profile.Username, profile.Password, profile.AvatarLink, profile.LastActivity, profile.CreatedAt, profile.CreatedFromIp, profile.DeletedAt, profile.IsActivate)
	return err
}

func (r ProfileRepository) Update(ctx context.Context, query string, args []interface{}) error {
	res, err := r.db.ExecContext(ctx, query, args...)
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
	res, err := r.db.ExecContext(ctx, `update profile set deleted_at = $1 where username = $2`, time.Now(), id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return errors.New("Zero rows affected ")
	}
	return err
}
