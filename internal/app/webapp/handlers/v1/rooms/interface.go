package rooms

import "github.com/gin-gonic/gin"

type RoomHandler interface {
	GetRoom(c *gin.Context)
	ListRooms(c *gin.Context)

	GetRoomReservations(c *gin.Context)
	GetReservation(c *gin.Context)

	CreateReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
}
