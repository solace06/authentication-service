package user

import (
	"github.com/uptrace/bun"
)

var s *Scope

type Scope struct {
	db *bun.DB
}

func NewScope(db *bun.DB) error {
	s = &Scope{
		db: db,
	}

	return nil
}
