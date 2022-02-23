package service

import (
	"context"

	"github.com/sirupsen/logrus"

	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

func (s *serviceImpl) GetRoom(ctx context.Context, id int) (*api.Room, error) {
	room, err := s.db.Rooms.GetRoom(ctx, id)
	if err != nil {
		logrus.Errorf("Error getting room[ID:%d]: %v", id, err)
		return nil, ErrDBError
	}
	if room == nil {
		return nil, ErrNotFound
	}
	return s.converter.ToAPIRoom(room), nil
}
