package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"

	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	api "github.com/FurmanovD/roommgr/pkg/api/v1"
	"github.com/FurmanovD/roommgr/test/mocks"
)

var defaultTestServiceConfig = Config{}

func TestGetReservation(t *testing.T) {

	testCases := map[string]interface{}{
		"ErrorDBQuery":  caseGetReservationErrorDBQuery(t),
		"ErrorNotFound": caseGetReservationErrorNotFound(t),
		"OK":            caseGetReservationOK(t),
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if err := quick.Check(tc, nil); err != nil {
				t.Errorf("%v case failed with an error: %+v", name, err)
			}
		})
	}
}

func caseGetReservationErrorDBQuery(t *testing.T) interface{} {
	return func(id int, dbErrorStr string) bool {
		dbError := errors.New(dbErrorStr)
		if dbErrorStr == "" {
			dbError = errors.New("non-empty error")
		}

		// create a station service:
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRoomRepository := mocks.NewMockRoomRepository(mockCtrl)
		mockRoomRepository.EXPECT().GetReservation(gomock.Any(), nil, id).Return(nil, dbError).Times(1)

		repo := repository.NewRepository(nil)
		repo.Rooms = mockRoomRepository

		// Test function
		_, err := NewService(defaultTestServiceConfig, repo, nil).GetReservation(context.Background(), id)

		return assert.Equal(t, ErrDBError, err)
	}
}

func caseGetReservationErrorNotFound(t *testing.T) interface{} {
	return func(id int) bool {

		// create a station service:
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRoomRepository := mocks.NewMockRoomRepository(mockCtrl)
		mockRoomRepository.EXPECT().GetReservation(gomock.Any(), nil, id).
			Return(nil, nil).
			Times(1)

		repo := repository.NewRepository(nil)
		repo.Rooms = mockRoomRepository

		// Test function
		_, err := NewService(defaultTestServiceConfig, repo, nil).GetReservation(context.Background(), id)

		return assert.Equal(t, ErrNotFound, err)
	}
}

func caseGetReservationOK(t *testing.T) interface{} {
	return func(reservationJoin repository.ReservationJoinUserJoinRoom) bool {

		// create a station service:
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockRoomRepository := mocks.NewMockRoomRepository(mockCtrl)
		mockRoomRepository.EXPECT().GetReservation(gomock.Any(), nil, reservationJoin.Reservation.ID).
			Return(&reservationJoin, nil).
			Times(1)

		repo := repository.NewRepository(nil)
		repo.Rooms = mockRoomRepository

		testAPIReservation := api.Reservation{
			ID: reservationJoin.Reservation.ID,
		}
		mockConverter := mocks.NewMockAPIDBConverter(mockCtrl)
		mockConverter.EXPECT().ToAPIReservation(&reservationJoin).
			Return(&testAPIReservation).
			Times(1)

		// Test function
		retObj, err := NewService(defaultTestServiceConfig, repo, mockConverter).
			GetReservation(context.Background(), reservationJoin.Reservation.ID)

		return assert.NoError(t, err) &&
			assert.True(t, reflect.DeepEqual(*retObj, testAPIReservation), "expected and returned objcets do not match")
	}
}
