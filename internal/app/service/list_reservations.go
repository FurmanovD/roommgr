package service

import (
	"context"
	"time"

	api "github.com/FurmanovD/roommgr/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

func (s *serviceImpl) ListRoomReservations(ctx context.Context, roomID int) (api.Reservations, error) {
	reservations, err := s.db.Rooms.GetRoomReservations(
		ctx,
		nil,
		roomID,
		time.Time{},
	)
	if err != nil {
		logrus.Error("Error listing reservations: %v", err)
		return nil, ErrDBError
	}
	return s.converter.ToAPIReservations(reservations), nil
}
