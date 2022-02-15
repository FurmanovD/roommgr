package rooms

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"

	"github.com/FurmanovD/roommgr/internal/app/service"
	v1 "github.com/FurmanovD/roommgr/internal/app/webapp/handlers/v1"
	"github.com/FurmanovD/roommgr/internal/app/webapp/weberror"
	"github.com/FurmanovD/roommgr/pkg/api/v1"
)

// roomHandler an EndpointHandler interface holder
type roomHandlerImpl struct {
	service service.RoomManagerService
}

// NewRoomHandler instantiates a HealthChecker
func NewRoomHandler(svc service.RoomManagerService) RoomHandler {
	return &roomHandlerImpl{
		service: svc,
	}
}

// ListRooms is a GET /api/v1/rooms handler
func (h *roomHandlerImpl) ListRooms(c *gin.Context) {

	rooms, err := h.service.ListRooms(c.Request.Context())
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	//return rooms array
	c.JSON(
		http.StatusOK,
		api.GetRoomsResponse(rooms),
	)
}

// GetRoom is a GET /api/v1/rooms/:roomID handler
func (h *roomHandlerImpl) GetRoom(c *gin.Context) {

	roomID, err := strconv.Atoi(c.Param(v1.KeyRoomID))
	if err != nil || roomID == 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				v1.KeyRoomID+" expected to be a natural number",
			),
		)
		return
	}

	room, err := h.service.GetRoom(c.Request.Context(), roomID)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	//return room object
	c.JSON(
		http.StatusOK,
		room,
	)
}

// GetRoomReservations is a GET /api/v1/rooms/:roomID/reservations handler
func (h *roomHandlerImpl) GetRoomReservations(c *gin.Context) {

	roomID, err := strconv.Atoi(c.Param(v1.KeyRoomID))
	if err != nil || roomID == 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				v1.KeyRoomID+" expected to be a natural number",
			),
		)
		return
	}

	reservations, err := h.service.ListRoomReservations(c.Request.Context(), roomID)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	//return room object
	c.JSON(
		http.StatusOK,
		api.ReservationsResponse(reservations),
	)
}

// GetReservation is a GET /api/v1/rooms/-/reservations/:reservationID handler
func (h *roomHandlerImpl) GetReservation(c *gin.Context) {

	reservationID, err := strconv.Atoi(c.Param(v1.KeyReservationID))
	if err != nil || reservationID == 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				v1.KeyReservationID+" expected to be a natural number",
			),
		)
		return
	}

	reservation, err := h.service.GetReservation(c.Request.Context(), reservationID)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	//return room object
	c.JSON(
		http.StatusOK,
		reservation,
	)
}

// CreateReservation is a POST /api/v1/rooms/-/reservations/ handler
func (h *roomHandlerImpl) CreateReservation(c *gin.Context) {

	request, err := parseAndValidateCreateReservationRequest(c)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				err.Error(),
			),
		)
		return
	}

	reservation, err := h.service.CreateReservation(
		c.Request.Context(),
		request.RoomID,
		request.UserToken, // UserID
		request.StartTime,
		request.StartTime.Add(time.Hour), // 1 hour period by default
	)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(err, ""),
		)
		return
	}

	//return reservation object
	c.JSON(
		http.StatusOK,
		reservation,
	)
}

// CreateReservation is a DELETE /api/v1/rooms/-/reservations/:id handler
func (h *roomHandlerImpl) DeleteReservation(c *gin.Context) {

	reservationID, err := strconv.Atoi(c.Param(v1.KeyReservationID))
	if err != nil || reservationID == 0 {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				v1.KeyReservationID+" expected to be a natural number",
			),
		)
		return
	}

	request, err := parseAndValidateDeleteReservationRequest(c)
	if err != nil {
		c.JSON(
			weberror.GetWebResponse(
				service.ErrInvalidRequest,
				err.Error(),
			),
		)
		return
	}

	err = h.service.DeleteReservation(
		c.Request.Context(),
		request.UserToken, // UserID
		reservationID,
	)

	c.JSON(
		weberror.GetWebResponse(err, ""),
	)
}

func parseAndValidateCreateReservationRequest(c *gin.Context) (*api.CreateReservationRequest, error) {

	request := api.CreateReservationRequest{}
	if err := c.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("invalid POST data structure: %v", err)
	}

	v := validate.Struct(request)
	if !v.Validate() {
		return nil, fmt.Errorf(v.Errors.Error())
	}

	// validate start time
	startTimeStr := request.StartTime.UTC().Format(api.TimeFormat)
	// "2006-01-02T15:04:05"
	if len(startTimeStr) < len(api.TimeFormat) ||
		startTimeStr[14:] != "00:00" {
		return nil, fmt.Errorf("start_time should not contain minutes and seconds yet")
	}
	return &request, nil
}

func parseAndValidateDeleteReservationRequest(c *gin.Context) (*api.DeleteReservationRequest, error) {

	request := api.DeleteReservationRequest{}
	if err := c.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("invalid POST data structure: %v", err)
	}

	v := validate.Struct(request)
	if !v.Validate() {
		return nil, fmt.Errorf(v.Errors.Error())
	}
	return &request, nil
}
