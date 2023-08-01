package types

import "github.com/uptrace/bun"

type Hotel struct {
	bun.BaseModel `bun:"table:hotels,alias:h"`
	HotelID       int64  `bun:",pk,autoincrement" json:"hotelID"`
	Name          string `json:"name"`
	Location      string `json:"location"`
	Rating        uint8  `json:"rating"`
}
