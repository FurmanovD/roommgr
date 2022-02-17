// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/db/repository/room_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	automodel "github.com/FurmanovD/roommgr/internal/pkg/db/automodel"
	repository "github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockRoomRepository is a mock of RoomRepository interface.
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository.
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance.
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// DeleteReservation mocks base method.
func (m *MockRoomRepository) DeleteReservation(ctx context.Context, executor repository.Transaction, reservationID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReservation", ctx, executor, reservationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReservation indicates an expected call of DeleteReservation.
func (mr *MockRoomRepositoryMockRecorder) DeleteReservation(ctx, executor, reservationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReservation", reflect.TypeOf((*MockRoomRepository)(nil).DeleteReservation), ctx, executor, reservationID)
}

// GetDBReservations mocks base method.
func (m *MockRoomRepository) GetDBReservations(ctx context.Context, executor repository.Transaction, id, userID, roomID int, startingFrom time.Time) (automodel.ReservationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBReservations", ctx, executor, id, userID, roomID, startingFrom)
	ret0, _ := ret[0].(automodel.ReservationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBReservations indicates an expected call of GetDBReservations.
func (mr *MockRoomRepositoryMockRecorder) GetDBReservations(ctx, executor, id, userID, roomID, startingFrom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBReservations", reflect.TypeOf((*MockRoomRepository)(nil).GetDBReservations), ctx, executor, id, userID, roomID, startingFrom)
}

// GetReservation mocks base method.
func (m *MockRoomRepository) GetReservation(ctx context.Context, executor repository.Transaction, id int) (*repository.ReservationJoinUserJoinRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservation", ctx, executor, id)
	ret0, _ := ret[0].(*repository.ReservationJoinUserJoinRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservation indicates an expected call of GetReservation.
func (mr *MockRoomRepositoryMockRecorder) GetReservation(ctx, executor, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservation", reflect.TypeOf((*MockRoomRepository)(nil).GetReservation), ctx, executor, id)
}

// GetRoom mocks base method.
func (m *MockRoomRepository) GetRoom(ctx context.Context, id int) (*automodel.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, id)
	ret0, _ := ret[0].(*automodel.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomRepositoryMockRecorder) GetRoom(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomRepository)(nil).GetRoom), ctx, id)
}

// GetRoomReservations mocks base method.
func (m *MockRoomRepository) GetRoomReservations(ctx context.Context, executor repository.Transaction, roomID int, startFrom time.Time) ([]*repository.ReservationJoinUserJoinRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomReservations", ctx, executor, roomID, startFrom)
	ret0, _ := ret[0].([]*repository.ReservationJoinUserJoinRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomReservations indicates an expected call of GetRoomReservations.
func (mr *MockRoomRepositoryMockRecorder) GetRoomReservations(ctx, executor, roomID, startFrom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomReservations", reflect.TypeOf((*MockRoomRepository)(nil).GetRoomReservations), ctx, executor, roomID, startFrom)
}

// GetRooms mocks base method.
func (m *MockRoomRepository) GetRooms(ctx context.Context, filter repository.RoomFilter) (automodel.RoomSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRooms", ctx, filter)
	ret0, _ := ret[0].(automodel.RoomSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRooms indicates an expected call of GetRooms.
func (mr *MockRoomRepositoryMockRecorder) GetRooms(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRooms", reflect.TypeOf((*MockRoomRepository)(nil).GetRooms), ctx, filter)
}

// InsertReservation mocks base method.
func (m *MockRoomRepository) InsertReservation(ctx context.Context, executor repository.Transaction, reservation *automodel.Reservation) (*automodel.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertReservation", ctx, executor, reservation)
	ret0, _ := ret[0].(*automodel.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertReservation indicates an expected call of InsertReservation.
func (mr *MockRoomRepositoryMockRecorder) InsertReservation(ctx, executor, reservation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertReservation", reflect.TypeOf((*MockRoomRepository)(nil).InsertReservation), ctx, executor, reservation)
}
