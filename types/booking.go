package types

import (
	"time"

	"github.com/uptrace/bun"
)

type Booking struct {
	bun.BaseModel `bun:"table:bookings,alias:b"`
	BookingID     int64     `bun:",pk,autoincrement,unique" json:"bookingID"`
	UserID        int64     `json:"userID"`
	RoomID        int64     `json:"roomID"`
	CheckIn       time.Time `json:"checkIn"`  // ISO 8601 format
	CheckOut      time.Time `json:"checkOut"` // ISO 8601 format
	Persons       int       `json:"persons"`
	Cancelled     bool      `json:"cancelled"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
