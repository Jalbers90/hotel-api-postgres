package db

import (
	"context"

	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/uptrace/bun"
)

type BookingStore interface {
	GetHotelBookings(context.Context, int64) ([]*types.Booking, error)
	GetUserBookings(context.Context, int64) ([]*types.Booking, error)
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	CancelBooking(context.Context, int64) (*types.Booking, error)
}

type PGBookingStore struct {
	db *bun.DB
}

func NewPGBookingStore(db *bun.DB) *PGBookingStore {
	return &PGBookingStore{
		db: db,
	}
}

func (s *PGBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	var insertedBooking types.Booking
	_, err := s.db.NewInsert().Model(booking).Returning("*").Exec(ctx, &insertedBooking)
	if err != nil {
		return nil, err
	}
	return &insertedBooking, nil
}

func (s *PGBookingStore) GetUserBookings(ctx context.Context, userID int64) ([]*types.Booking, error) {
	var bookings []*types.Booking
	err := s.db.NewSelect().Model(&types.Booking{}).Where("user_id = ? AND cancelled = false", userID).Scan(ctx, &bookings)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *PGBookingStore) CancelBooking(ctx context.Context, bookingID int64) (*types.Booking, error) {
	booking := types.Booking{}
	_, err := s.db.NewUpdate().Model(&booking).Set("cancelled = ?", true).Where("booking_id = ?", bookingID).Exec(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.db.NewSelect().Model(&booking).Where("booking_id = ?", bookingID).Exec(ctx, &booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (s *PGBookingStore) GetHotelBookings(ctx context.Context, hotelID int64) ([]*types.Booking, error) {

	return nil, nil
}
