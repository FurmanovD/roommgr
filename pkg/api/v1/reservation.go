package api

import (
	"time"
)

type Reservations []*Reservation

type ReservationsResponse = Reservations

type Reservation struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	Room      string    `json:"room"`
	Company   string    `json:"company"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type CreateReservationRequest struct {
	// TODO: move UserToken to URL parameters
	// actually, it is just a UserID until authentication and token(JWT?) processing is added
	UserToken int       `json:"user_token" validate:"required|min:1"`
	RoomID    int       `json:"room_id" validate:"required|min:1"`
	StartTime time.Time `json:"start_time" validate:"required"`
	// EndTime   time.Time `json:"end_time"`
}

type DeleteReservationRequest struct {
	// TODO: move UserToken to URL parameters
	// actually, it is just a UserID until authentication and token(JWT?) processing is added
	UserToken int `json:"user_token" validate:"required|min:1"`
}
