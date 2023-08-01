package api

import "github.com/Jalbers90/hotel-api-postgres/db"

type RoomHandler struct {
	store db.RoomStore
}

func NewRoomHandler(store db.RoomStore) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}
