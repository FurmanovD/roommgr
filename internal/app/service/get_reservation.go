package service

import (
	"context"

	"github.com/sirupsen/logrus"

	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

func (s *serviceImpl) GetReservation(ctx context.Context, id int) (*api.Reservation, error) {
	reservation, err := s.db.Rooms.GetReservation(ctx, nil, id)
	if err != nil {
		logrus.Errorf("Error getting reservation[ID:%d]: %v", id, err)
		return nil, ErrDBError
	}
	return s.converter.ToAPIReservation(reservation), nil
}
