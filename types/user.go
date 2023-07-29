package types

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel     `bun:"table:users,alias:u"`
	ID                int64  `bun:",pk,autoincrement" json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
	Admin             bool   `json:"admin"`
}
