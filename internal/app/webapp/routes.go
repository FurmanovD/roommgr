package webapp

import (
	"github.com/FurmanovD/roommgr/internal/app/webapp/handlers/healthcheck"
	v1 "github.com/FurmanovD/roommgr/internal/app/webapp/handlers/v1"
	"github.com/FurmanovD/roommgr/internal/app/webapp/handlers/v1/rooms"
)

const (
	PathHealthCheck = "healthcheck"

	PathCompanies        = "companies"
	PathUsers            = "users"
	PathRooms            = "rooms"
	PathRoom             = "rooms/:" + v1.KeyRoomID
	PathRoomReservations = "rooms/:" + v1.KeyRoomID + "/reservations"
	PathReservation      = "rooms/-/reservations/:" + v1.KeyReservationID
	PathAllReservations  = "rooms/-/reservations/"
)

// RegisterRoutes registers all the routes
func (s *webServer) RegisterRoutes() {

	s.registerRootRoutes()

	s.registerAPIv1Routes()

	// add all routes here
}

func (s *webServer) registerRootRoutes() {

	// add healthcheck url handler
	s.routerGroupRoot.GET(PathHealthCheck, healthcheck.NewHealthChecker(s.roommgrService).GetHealth)
}

func (s *webServer) registerAPIv1Routes() {

	roomsHandler := rooms.NewRoomHandler(s.roommgrService)

	s.registerAPIv1RoomRoutes(roomsHandler)

	s.registerAPIv1ReservationRoutes(roomsHandler)

	// add all API v1 routes here
}

func (s *webServer) registerAPIv1RoomRoutes(handler rooms.RoomHandler) {

	s.routerGroupAPIV1.GET(PathRooms, handler.ListRooms)
	s.routerGroupAPIV1.GET(PathRoom, handler.GetRoom)
}

func (s *webServer) registerAPIv1ReservationRoutes(handler rooms.RoomHandler) {

	s.routerGroupAPIV1.GET(PathRoomReservations, handler.GetRoomReservations)
	s.routerGroupAPIV1.GET(PathReservation, handler.GetReservation)

	s.routerGroupAPIV1.POST(PathAllReservations, handler.CreateReservation)

	s.routerGroupAPIV1.DELETE(PathReservation, handler.DeleteReservation)
}
