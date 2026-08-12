package user

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `bun:"id,pk,autoincrement" json:"id"`
	Name         string    `bun:"name" json:"name"`
	Email        string    `bun:"email" json:"email"`
	Role         string    `bun:"role" json:"role"`
	PasswordHash string    `bun:"password_hash" json:"-"`
	CreatedAt    time.Time `bun:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bun:"updated_at" json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}
