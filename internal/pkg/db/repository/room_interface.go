package repository

import (
	"context"
	"time"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
)

// RoomRepository contains all functions required to manage Room objects and their state
type RoomRepository interface {
	GetRoom(ctx context.Context, id int) (*automodel.Room, error)
	GetRooms(ctx context.Context, filter RoomFilter) (automodel.RoomSlice, error)

	GetReservation(
		ctx context.Context,
		executor Transaction,
		id int,
	) (*ReservationJoinUserJoinRoom, error)
	GetRoomReservations(
		ctx context.Context,
		executor Transaction,
		roomID int,
		startFrom time.Time,
	) ([]*ReservationJoinUserJoinRoom, error)

	GetDBReservations(
		ctx context.Context,
		executor Transaction,
		ID int,
		userID int,
		roomID int,
		startingFrom time.Time,
	) (automodel.ReservationSlice, error)

	InsertReservation(
		ctx context.Context,
		executor Transaction,
		reservation *automodel.Reservation,
	) (*automodel.Reservation, error)

	DeleteReservation(
		ctx context.Context,
		executor Transaction,
		reservationID int,
	) error
}
