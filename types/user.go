package types

import (
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel     `bun:"table:users,alias:u"`
	ID                int64  `bun:",pk,autoincrement" json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
	Admin             bool   `json:"-"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func CreateUserFromParams(params CreateUserParams) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12) // use CompareHashAndPassword to compare encryped
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}
