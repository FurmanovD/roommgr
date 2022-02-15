package repository

import (
	"github.com/FurmanovD/roommgr/pkg/sqldb"
)

type Repository struct {
	TxCreator TxCreator
	Rooms     RoomRepository
	Users     UserRepository
}

func NewRepository(db sqldb.SqlDB) *Repository {
	return &Repository{
		TxCreator: NewTxCreator(db),
		Rooms:     NewRoomRepository(db),
		Users:     NewUserRepository(db),
	}
}
