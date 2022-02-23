package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

func (s *serviceImpl) ListRoomReservations(ctx context.Context, roomID int) (api.Reservations, error) {
	reservations, err := s.db.Rooms.GetRoomReservations(
		ctx,
		nil,
		roomID,
		time.Time{},
	)
	if err != nil {
		logrus.Errorf("Error listing reservations: %v", err)
		return nil, ErrDBError
	}
	return s.converter.ToAPIReservations(reservations), nil
}
