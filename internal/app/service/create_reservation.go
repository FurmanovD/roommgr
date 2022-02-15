package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

func (s *serviceImpl) CreateReservation(
	ctx context.Context,
	roomID int,
	userID int,
	startTime time.Time,
	endTime time.Time,
) (*api.Reservation, error) {

	tx, err := s.db.TxCreator.CreateTransaction(ctx)
	if err != nil {
		logrus.Error("error creating transaction")
		return nil, ErrDBError
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	checkTimeStart := time.Now().UTC()
	checkTimeStart = checkTimeStart.Add(-24 * time.Hour) // to skip old reservations
	if startTime.Before(checkTimeStart) {
		logrus.Error("cannot create reservation with start time more than 24 hours ago")
		return nil, ErrCannotCreate
	}

	// check if reservation is possible:
	dbReservations, err := s.db.Rooms.GetDBReservations(
		ctx,
		tx,
		-1, // ignore id
		-1, // ignore user
		roomID,
		checkTimeStart,
	)
	if err != nil {
		logrus.Errorf("reading room[ID:%d] reservations failed: %w", roomID, err)
		return nil, ErrDBError
	}

	// Check overlapping
	var overlap *automodel.Reservation
	for _, item := range dbReservations {
		if item.StartTime.Equal(startTime) {
			overlap = item
			break
		}
		if item.StartTime.Before(startTime) &&
			(item.EndTime.Equal(endTime) || item.EndTime.After(endTime)) {
			overlap = item
			break
		}
	}
	if overlap != nil {
		logrus.Errorf(
			"reservation to create[roomID:%d,start:%v;end%v] overlaps with another[ID:%d,start:%v;end%v]",
			roomID,
			startTime,
			endTime,
			overlap.ID,
			overlap.StartTime,
			overlap.EndTime,
		)
		return nil, ErrCannotCreate
	}

	//
	reservation, err := s.db.Rooms.InsertReservation(
		ctx,
		tx,
		&automodel.Reservation{
			UserID:    userID,
			RoomID:    roomID,
			StartTime: startTime,
			EndTime:   endTime,
		},
	)

	// Get Reservation to return:
	reservationJoin, err := s.db.Rooms.GetReservation(ctx, tx, reservation.ID)
	if err != nil {
		logrus.Errorf(
			"error adding new reservation[%+v] to DB: %v",
			reservation,
			err,
		)
		return nil, ErrDBError
	}

	if err = tx.Commit(); err != nil {
		logrus.Errorf(
			"failed to commit new reservation[%+v] to DB: %v",
			reservation,
			err,
		)
		return nil, ErrDBError
	}
	if reservationJoin == nil {
		logrus.Errorf(
			"failed to read recently created reservation[%+v] to return",
			reservation,
		)
		return nil, ErrInternalServerError
	}

	return s.converter.ToAPIReservation(reservationJoin), nil
}
