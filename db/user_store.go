package db

import (
	"context"

	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/uptrace/bun"
)

const USERTABLE = "users"

type UserStore interface {
	GetUsers(context.Context, types.Map) ([]*types.User, error)
	GetUserByID(context.Context, int) (*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
}

type PGUserStore struct {
	db *bun.DB
}

func NewPGUserStore(db *bun.DB) *PGUserStore {
	return &PGUserStore{
		db: db,
	}
}

func (s *PGUserStore) GetUsers(ctx context.Context, m types.Map) ([]*types.User, error) {
	// query db and return all users
	var users []*types.User
	err := s.db.NewSelect().Model(&types.User{}).Scan(ctx, &users)
	if err != nil {
		return []*types.User{}, err
	}
	return users, nil
}

func (s *PGUserStore) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	var user types.User
	err := s.db.NewSelect().Model(&types.User{}).Where("id = ?", id).Scan(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PGUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	err := s.db.NewSelect().Model(&types.User{}).Where("email = ?", email).Scan(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PGUserStore) InsertUser(ctx context.Context, userToInsert *types.User) (*types.User, error) {
	var newUser types.User
	_, err := s.db.NewInsert().Model(userToInsert).Returning("*").Exec(ctx, &newUser) // would not actually want to return password
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
