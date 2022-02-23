package api

import (
	"time"
)

type GetRoomsRequest struct{}

type GetRoomsResponse = Rooms

type Rooms []*Room

type Room struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
