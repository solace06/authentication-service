package user

import (
	"context"
	"fmt"
	"time"

	"github.com/solace06/auth-service/pkg"
)

func (s *Scope) CreateUser(ctx context.Context, req CreateUserRequest) error {

	var user *User

	//check if the user already exists
	user, err := s.FetchUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("user already exists")
	}

	//generate password hash
	passwordHash, err := pkg.HashPassword(req.Password)
	if err != nil {
		return err
	}

	currTime := time.Now()

	user = &User{
		Name:         req.Name,
		Email:        req.Email,
		Role:         req.Role,
		PasswordHash: passwordHash,
		CreatedAt:    currTime,
		UpdatedAt:    currTime,
	}

	err = s.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
