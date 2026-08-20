package user

import (
	"context"
	"database/sql"
	"errors"
)

func (s *Scope) FetchUserByEmail(ctx context.Context, email string) (*User, error) {

	var user User

	err := s.db.NewSelect().Model(&user).Where("email = ?", email).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Scope) InsertUser(ctx context.Context, user *User) error {

	_, err := s.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
