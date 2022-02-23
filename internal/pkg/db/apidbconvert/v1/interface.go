package apidbconvert

import (
	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

type APIDBConverter interface {
	// these functions convert a DB structure(s) to API object(s) and vice versa
	ToAPIRoom(room *automodel.Room) *api.Room
	ToAPIRooms(rooms automodel.RoomSlice) api.Rooms

	ToDBRoom(room *api.Room) *automodel.Room
	ToDBRooms(rooms api.Rooms) automodel.RoomSlice

	ToAPIReservation(reservationJoin *repository.ReservationJoinUserJoinRoom) *api.Reservation
	ToAPIReservations(reservationJoins []*repository.ReservationJoinUserJoinRoom) api.Reservations
}
