package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *serviceImpl) DeleteReservation(
	ctx context.Context,
	userID int,
	reservationID int,
) error {

	// check if reservation is made by the same user TODO: or check privileges
	dbReservations, err := s.db.Rooms.GetDBReservations(
		ctx,
		nil,
		reservationID,
		-1, // ignore user
		-1, // ignore room
		time.Time{},
	)
	if err != nil {
		logrus.Errorf("reading reservation[ID:%d] failed: %w", reservationID, err)
		return ErrDBError
	}
	if len(dbReservations) == 0 {
		logrus.Errorf("no reservation[ID:%d] found", reservationID)
		return ErrNotFound
	}
	if dbReservations[0].UserID != userID {
		logrus.Errorf(
			"reservation[ID:%d] was created by user[ID:%d] but request to delete is from [ID:%d]",
			reservationID,
			dbReservations[0].UserID,
			userID,
		)
		return ErrObjAccessDenied
	}

	// Delete
	if err = s.db.Rooms.DeleteReservation(
		ctx,
		nil,
		reservationID,
	); err != nil {
		logrus.Errorf(
			"error deleting reservation[ID:%d]: %v",
			reservationID,
			err,
		)
		return ErrDBError
	}

	return nil
}
