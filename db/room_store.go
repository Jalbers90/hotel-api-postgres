package db

import (
	"context"

	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/uptrace/bun"
)

type RoomStore interface {
	GetRooms(context.Context, types.Map) ([]*types.Room, error)
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type PGRoomStore struct {
	db *bun.DB
}

func NewPGRoomStore(db *bun.DB) *PGRoomStore {
	return &PGRoomStore{
		db: db,
	}
}

func (s *PGRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	var newRoom types.Room
	_, err := s.db.NewInsert().Model(room).Returning("*").Exec(ctx, &newRoom)
	if err != nil {
		return nil, err
	}
	return &newRoom, nil
}
