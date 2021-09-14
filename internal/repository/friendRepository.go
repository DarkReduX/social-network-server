package repository

import (
	"context"
	"database/sql"
	"errors"
)

type FriendRepository struct {
	db *sql.DB
}

func NewFriendRepository(db *sql.DB) *FriendRepository {
	return &FriendRepository{db: db}
}

func (r FriendRepository) AddFriendRequest(ctx context.Context, userId string, friendId string) error {
	query := `select * from add_friend_request($1,$2)`
	res, err := r.db.ExecContext(ctx, query, userId, friendId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Couldn't add friend in db ")
	}

	return err
}

func (r FriendRepository) ProcessFriendRequest(ctx context.Context, userId string, friendId string, requestType string) error {
	query := `select * from process_friend_request($1,$2,$3)`
	res, err := r.db.ExecContext(ctx, query, userId, friendId, requestType)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Couldn't submit friend in db ")
	}

	return err
}
