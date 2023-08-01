package types

import "github.com/uptrace/bun"

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:r"`
	RoomID        int64   `bun:",pk,autoincrement" json:"roomID"`
	HotelID       int64   `json:"hotelID"`
	Kind          string  `json:"kind"` // "King", "Queen", "Suite", etc
	Price         float64 `json:"price"`
}
