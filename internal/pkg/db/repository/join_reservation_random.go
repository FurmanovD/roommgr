package repository

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/FurmanovD/go-kit/randomstring"
	"github.com/FurmanovD/go-kit/testing/testvalue"
	"github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
)

// Generate implements quick.Generator interface for an Item
func (ReservationJoinUserJoinRoom) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(
		ReservationJoinUserJoinRoom{}.GetRandom(r),
	)
}

func (ReservationJoinUserJoinRoom) GetRandom(r *rand.Rand) ReservationJoinUserJoinRoom {

	rnd := testvalue.NotNilRnd(r)

	now := time.Time{}.UTC()
	maxID := 999999999

	// TODO add other values generation when needed
	// TODO add random generators for all included structure
	return ReservationJoinUserJoinRoom{
		Reservation: automodel.Reservation{
			ID:        rnd.Intn(maxID),
			RoomID:    rnd.Intn(maxID),
			UserID:    rnd.Intn(maxID),
			StartTime: now.Add(2 * time.Hour),
			EndTime:   now.Add(3 * time.Hour),
		},

		Room: automodel.Room{
			Name:      null.StringFrom(randomstring.NonEmptyUTF8Printable(30, rnd)),
			CompanyID: rnd.Intn(maxID),
		},

		User: automodel.User{
			FirstName: randomstring.NonEmptyUTF8Printable(30, rnd),
			LastName:  randomstring.NonEmptyUTF8Printable(30, rnd),
			Email:     null.StringFrom(randomstring.NonEmptyUTF8Printable(30, rnd)),
		},

		Company: automodel.Company{
			Name: null.StringFrom(randomstring.NonEmptyUTF8Printable(30, rnd)),
		},
	}
}
