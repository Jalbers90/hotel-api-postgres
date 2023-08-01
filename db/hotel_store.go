package db

import (
	"context"

	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/uptrace/bun"
)

type HotelStore interface {
	GetHotels(context.Context, types.Map) ([]*types.Hotel, error)
	GetHotelByID(context.Context, int) (*types.Hotel, error)
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	GetHotelRooms(context.Context, int) ([]*types.Room, error)
}

type PGHotelStore struct {
	db *bun.DB
}

func NewPGHotelStore(db *bun.DB) *PGHotelStore {
	return &PGHotelStore{
		db: db,
	}
}

func (s *PGHotelStore) GetHotels(ctx context.Context, filters types.Map) ([]*types.Hotel, error) {
	var hotels []*types.Hotel
	err := s.db.NewSelect().Model(&types.Hotel{}).Scan(ctx, &hotels)
	if err != nil {
		return []*types.Hotel{}, err
	}
	return hotels, nil
}

func (s *PGHotelStore) GetHotelByID(ctx context.Context, id int) (*types.Hotel, error) {
	var hotel types.Hotel
	err := s.db.NewSelect().Model(&types.Hotel{}).Where("hotel_id = ?", id).Scan(ctx, &hotel)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (s *PGHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	var newHotel types.Hotel
	_, err := s.db.NewInsert().Model(hotel).Returning("*").Exec(ctx, &newHotel)
	if err != nil {
		return nil, err
	}
	return &newHotel, nil
}

func (s *PGHotelStore) GetHotelRooms(ctx context.Context, hotelID int) ([]*types.Room, error) {
	var rooms []*types.Room
	err := s.db.NewSelect().Model(&types.Room{}).Where("hotel_id = ?", hotelID).Scan(ctx, &rooms)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
