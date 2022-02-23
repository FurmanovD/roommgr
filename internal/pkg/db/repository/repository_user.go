package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	"github.com/FurmanovD/roommgr/pkg/sqldb"
)

type userRepositoryImpl struct {
	db sqldb.SqlDB
}

func NewUserRepository(db sqldb.SqlDB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// ================= interface methods =================================================
func (r *userRepositoryImpl) GetUser(ctx context.Context, id int) (*automodel.User, error) {
	user, err := automodel.Users(
		automodel.UserWhere.ID.EQ(id),
	).One(ctx, r.db.Connection())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
