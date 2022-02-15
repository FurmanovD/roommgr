package service

import (
	"context"

	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

func (s *serviceImpl) ListRooms(ctx context.Context) (api.Rooms, error) {

	rooms, err := s.db.Rooms.GetRooms(ctx,
		repository.RoomFilter{
			IncludeDeleted: false,
		},
	)
	if err != nil {
		logrus.Error("Error getting rooms: %v", err)
		return nil, ErrDBError
	}
	return s.converter.ToAPIRooms(rooms), nil
}
