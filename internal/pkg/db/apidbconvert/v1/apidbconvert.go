package apidbconvert

import (
	"github.com/volatiletech/null/v8"

	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
)

type apiDBConverterImpl struct {
}

func NewAPIDBConverter() APIDBConverter {
	return &apiDBConverterImpl{}
}

func (c *apiDBConverterImpl) ToAPIRoom(room *automodel.Room) *api.Room {
	if room == nil {
		return nil
	}

	return &api.Room{
		ID:        room.ID,
		Name:      room.Name.String,
		CreatedAt: room.CreatedAt,
	}
}

func (c *apiDBConverterImpl) ToAPIRooms(rooms automodel.RoomSlice) api.Rooms {
	if rooms == nil {
		return nil
	}

	res := make(api.Rooms, len(rooms))
	for i, room := range rooms {
		res[i] = c.ToAPIRoom(room)
	}
	return res
}

func (c *apiDBConverterImpl) ToDBRoom(room *api.Room) *automodel.Room {
	if room == nil {
		return nil
	}

	var name *string
	if room.Name != "" {
		s := room.Name
		name = &s
	}
	return &automodel.Room{
		ID:        room.ID,
		Name:      null.StringFromPtr(name),
		CreatedAt: room.CreatedAt,
	}
}

func (c *apiDBConverterImpl) ToDBRooms(rooms api.Rooms) automodel.RoomSlice {
	if rooms == nil {
		return nil
	}

	res := make(automodel.RoomSlice, len(rooms))
	for i, room := range rooms {
		res[i] = c.ToDBRoom(room)
	}
	return res
}

func (c *apiDBConverterImpl) ToAPIReservation(
	reservationJoin *repository.ReservationJoinUserJoinRoom,
) *api.Reservation {
	if reservationJoin == nil {
		return nil
	}

	var name string
	if reservationJoin.User.FirstName != "" {
		name += reservationJoin.User.FirstName
	}
	if reservationJoin.User.LastName != "" {
		if name != "" {
			name += " "
		}
		name += reservationJoin.User.LastName
	}

	return &api.Reservation{
		ID:        reservationJoin.Reservation.ID,
		User:      name,
		Room:      reservationJoin.Room.Name.String,
		Company:   reservationJoin.Company.Name.String,
		StartTime: reservationJoin.Reservation.StartTime,
		EndTime:   reservationJoin.Reservation.EndTime,
	}
}

func (c *apiDBConverterImpl) ToAPIReservations(
	reservationJoins []*repository.ReservationJoinUserJoinRoom,
) api.Reservations {
	if reservationJoins == nil {
		return nil
	}

	res := make(api.Reservations, len(reservationJoins))
	for i, r := range reservationJoins {
		res[i] = c.ToAPIReservation(r)
	}
	return res
}
