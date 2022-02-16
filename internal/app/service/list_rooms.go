package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

func (s *serviceImpl) ListRooms(ctx context.Context) (api.Rooms, error) {
	rooms, err := s.db.Rooms.GetRooms(ctx,
		repository.RoomFilter{
			IncludeDeleted: false,
		},
	)
	if err != nil {
		logrus.Errorf("Error getting rooms: %v", err)
		return nil, ErrDBError
	}
	return s.converter.ToAPIRooms(rooms), nil
}
