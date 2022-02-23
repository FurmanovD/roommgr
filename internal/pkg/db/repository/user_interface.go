package repository

import (
	"context"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
)

// UserRepository contains all functions required to manage User objects and their state
type UserRepository interface {
	GetUser(ctx context.Context, id int) (*automodel.User, error)
}
