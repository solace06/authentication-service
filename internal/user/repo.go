package user

import "context"

func (s *Scope) FetchUserByEmail(ctx context.Context, email string) (*User, error) {

	var user User

	err := s.db.NewSelect().Model(&user).Where("email = ?", email).Scan(ctx)
	if err != nil{
		return nil, err
	}

	return &user, nil
}
