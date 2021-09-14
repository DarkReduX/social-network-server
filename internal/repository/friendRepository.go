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

func (r FriendRepository) Add(ctx context.Context, userId string, friendId string) error {
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

func (r FriendRepository) Delete(ctx context.Context, userId string, friendId string) error {
	query := `select * from delete_friend($1,$2)`
	res, err := r.db.ExecContext(ctx, query, userId, friendId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Couldn't delete friend in db ")
	}

	return err
}

func (r FriendRepository) SubmitFriend(ctx context.Context, userId string, friendId string) error {
	query := `select * from submit_friend_request($1,$2)`
	res, err := r.db.ExecContext(ctx, query, userId, friendId)
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
