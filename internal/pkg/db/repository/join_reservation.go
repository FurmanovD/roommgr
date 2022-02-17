package repository

import (
	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
)

// SQL Joins

type ReservationJoinUserJoinRoom struct {
	Reservation automodel.Reservation `boil:"reservations,bind"`
	User        automodel.User        `boil:"users,bind"`
	Room        automodel.Room        `boil:"rooms,bind"`
	Company     automodel.Company     `boil:"companies,bind"`
}
