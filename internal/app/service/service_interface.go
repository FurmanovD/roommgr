package service

import (
	"context"
	"time"

	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

type RoomManagerService interface {
	GetRoom(ctx context.Context, id int) (*api.Room, error)
	ListRooms(ctx context.Context) (api.Rooms, error)

	GetReservation(ctx context.Context, id int) (*api.Reservation, error)
	ListRoomReservations(ctx context.Context, roomID int) (api.Reservations, error)

	CreateReservation(
		ctx context.Context,
		roomID int,
		userID int,
		startTime time.Time,
		endTime time.Time,
	) (*api.Reservation, error)

	DeleteReservation(
		ctx context.Context,
		userID int,
		reservationID int,
	) error
}
