package api

import (
	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/gofiber/fiber/v2"
)

// handle routes, keep decoupled from db stores

type UserHandler struct {
	store *db.PGUserStore
}

func NewUserHandler(store *db.PGUserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers(c.Context(), map[string]any{})
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id := 1
	user, err := h.store.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
